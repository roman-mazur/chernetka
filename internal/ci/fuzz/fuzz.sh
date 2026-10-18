#!/usr/bin/env bash
# Fuzzes the editor for $FUZZTIME, and outputs found=true if a failing input is saved to $CORPUS.

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
