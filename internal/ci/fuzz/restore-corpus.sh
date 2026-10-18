#!/usr/bin/env bash
# Restores the fuzz cache of GOCACHE from the fuzz-corpus artifact of the latest run that saved it.

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
