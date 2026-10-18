#!/usr/bin/env bash
# Commits the failing input found in $CORPUS to a new branch, and opens a draft PR showing the failure.

branch="fuzz/$GITHUB_RUN_ID"
run_url="$GITHUB_SERVER_URL/$GITHUB_REPOSITORY/actions/runs/$GITHUB_RUN_ID"
git config user.name 'github-actions[bot]'
git config user.email '41898282+github-actions[bot]@users.noreply.github.com'
git switch -c "$branch"
git add -- "$CORPUS"
git commit -m "internal/editor: add the failing input found by fuzzing" \
  -m "Found by the weekly fuzz run $run_url."
git push origin "$branch"
{
  echo "The weekly fuzzing found a failing input, committed to \`$CORPUS\`."
  echo
  echo "Run: $run_url"
  echo
  echo '```'
  grep -A40 -m1 -e '--- FAIL' fuzz.log || tail -n 60 fuzz.log
  echo '```'
} > pr-body.md
gh pr create --draft --base main --head "$branch" \
  --title "internal/editor: fuzzing failure from run $GITHUB_RUN_ID" \
  --body-file pr-body.md
