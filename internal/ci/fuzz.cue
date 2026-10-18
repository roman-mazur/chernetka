@extern(embed)

package ci

import (
	"list"

	"cue.dev/x/githubactions"
)

// The fuzz workflow fuzzes the editor weekly. A failing input is committed to a new branch with
// a draft PR, rather than to main: the seed corpus runs with go test, so it fails the tests until fixed.
// The scripts of the steps are in the fuzz directory.
workflows: fuzz: githubactions.#Workflow & {
	name: "Fuzz"
	on: {
		schedule: [{cron: "17 1 * * 6"}] // the night between Friday and Saturday, 01:17 UTC
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
			run: _ @embed(file=fuzz/check.sh, type=text)
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
				run: _ @embed(file=fuzz/restore-corpus.sh, type=text)
			},
			{
				name: "Fuzz"
				id:   "fuzz"
				run:  _ @embed(file=fuzz/fuzz.sh, type=text)
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
				run: _ @embed(file=fuzz/commit-failure.sh, type=text)
			},
		]])
	}
}
