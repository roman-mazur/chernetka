package ci

import "list"

// The fuzz workflow fuzzes the editor weekly. A failing input is committed to a new branch with
// a draft PR, rather than to main: the seed corpus runs with go test, so it fails the tests until fixed.
workflows: fuzz: #Workflow & {
	name: "Fuzz"
	on: {
		schedule: [{cron: "17 3 * * 1"}] // Mondays, 03:17 UTC
		workflow_dispatch: inputs: fuzztime: {
			description: "How long to fuzz, as -fuzztime"
			default:     "30m"
		}
	}

	// A scheduled run is skipped when main has not changed since the last successful run.
	// The manual runs always fuzz.
	jobs: check: {
		"runs-on": "ubuntu-latest"
		permissions: actions: "read"
		outputs: changed:     "${{ steps.last.outputs.changed }}"
		steps: [{
			name: "Compare with the last run"
			id:   "last"
			env: GH_TOKEN: "${{ github.token }}"
			run: #"""
				last=$(gh run list --repo "$GITHUB_REPOSITORY" --workflow fuzz.yml --branch main \
				  --status success --limit 1 --json headSha --jq '.[0].headSha // ""')
				echo "last fuzzed: ${last:-none}, main: $GITHUB_SHA"
				if [ "$GITHUB_EVENT_NAME" = schedule ] && [ "$last" = "$GITHUB_SHA" ]; then
				  echo "changed=false" >> "$GITHUB_OUTPUT"
				else
				  echo "changed=true" >> "$GITHUB_OUTPUT"
				fi
				"""#
		}]
	}

	jobs: fuzz: {
		needs:     "check"
		if:        "needs.check.outputs.changed == 'true'"
		"runs-on": "ubuntu-latest"
		permissions: {
			actions:         "read" // to download the corpus of the previous runs
			contents:        "write"
			"pull-requests": "write"
		}
		env: {
			FUZZTIME: "${{ inputs.fuzztime || '30m' }}"
			CORPUS:   "internal/editor/testdata/fuzz/FuzzEditor"
		}
		steps: list.Concat([#goSetup, [
			// The interesting inputs the fuzzing generates are kept in the fuzz cache of GOCACHE.
			// setup-go doesn't keep them: its cache is saved only when go.sum changes. They are kept
			// as an artifact instead: the actions cache is evicted when unused for 7 days,
			// which a weekly run, skipped when main doesn't change, would hit.
			{
				name: "Restore the corpus"
				env: GH_TOKEN: "${{ github.token }}"
				run: #"""
					dir="$(go env GOCACHE)/fuzz"
					echo "FUZZCACHE=$dir" >> "$GITHUB_ENV"
					run=$(gh api "repos/$GITHUB_REPOSITORY/actions/artifacts?name=fuzz-corpus&per_page=1" \
					  --jq '.artifacts[] | select(.expired | not) | .workflow_run.id')
					if [ -z "$run" ]; then
					  echo "no corpus of the previous runs, starting from the seeds"
					  exit 0
					fi
					gh run download "$run" --repo "$GITHUB_REPOSITORY" --name fuzz-corpus --dir "$dir"
					echo "restored $(find "$dir" -type f | wc -l) inputs from run $run"
					"""#
			},
			{
				name: "Fuzz"
				id:   "fuzz"
				run: #"""
					# -timeout is above the fuzzing time: the default 10m would stop a longer run.
					set +e
					go test -run '^$' -fuzz FuzzEditor -fuzztime "$FUZZTIME" -timeout 2h ./internal/editor 2>&1 | tee fuzz.log
					code=${PIPESTATUS[0]}
					set -e
					new=$(git status --porcelain --untracked-files=all -- "$CORPUS")
					if [ -n "$new" ]; then
					  echo "found=true" >> "$GITHUB_OUTPUT"
					elif [ "$code" -ne 0 ]; then
					  echo "fuzzing failed without a new input" >&2
					  exit "$code"
					fi
					"""#
			},
			{
				name: "Upload the log"
				if:   "always()"
				uses: "actions/upload-artifact@v4"
				with: {
					name: "fuzz-log"
					path: "fuzz.log"
				}
			},
			// Kept on a failure too: the inputs found before it are still interesting.
			{
				name: "Save the corpus"
				if:   "always() && steps.fuzz.outcome != 'skipped'"
				uses: "actions/upload-artifact@v4"
				with: {
					name:                "fuzz-corpus"
					path:                "${{ env.FUZZCACHE }}"
					"retention-days":    90
					"if-no-files-found": "ignore"
				}
			},
			{
				name: "Commit the failing input"
				if:   "steps.fuzz.outputs.found == 'true'"
				env: GH_TOKEN: "${{ github.token }}"
				run: #"""
					branch="fuzz/${{ github.run_id }}"
					git config user.name 'github-actions[bot]'
					git config user.email '41898282+github-actions[bot]@users.noreply.github.com'
					git switch -c "$branch"
					git add -- "$CORPUS"
					git commit -m "internal/editor: add the failing input found by fuzzing" \
					  -m "Found by the weekly fuzz run ${{ github.server_url }}/${{ github.repository }}/actions/runs/${{ github.run_id }}."
					git push origin "$branch"
					{
					  echo "The weekly fuzzing found a failing input, committed to \`$CORPUS\`."
					  echo
					  echo "Run: ${{ github.server_url }}/${{ github.repository }}/actions/runs/${{ github.run_id }}"
					  echo
					  echo '```'
					  grep -A40 -m1 -e '--- FAIL' fuzz.log || tail -n 60 fuzz.log
					  echo '```'
					} > pr-body.md
					gh pr create --draft --base main --head "$branch" \
					  --title "internal/editor: fuzzing failure from run ${{ github.run_id }}" \
					  --body-file pr-body.md
					"""#
			},
		]])
	}
}
