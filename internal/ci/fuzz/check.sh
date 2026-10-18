#!/usr/bin/env bash
# Tells if main has changed since the last successful fuzz run: a scheduled run is skipped otherwise.
# The manual runs always fuzz.

last=$(gh run list --repo "$GITHUB_REPOSITORY" --workflow fuzz.yml --branch main \
  --status success --limit 1 --json headSha --jq '.[0].headSha // ""')
echo "last fuzzed: ${last:-none}, main: $GITHUB_SHA"
if [ "$GITHUB_EVENT_NAME" = schedule ] && [ "$last" = "$GITHUB_SHA" ]; then
  echo "changed=false" >> "$GITHUB_OUTPUT"
else
  echo "changed=true" >> "$GITHUB_OUTPUT"
fi
