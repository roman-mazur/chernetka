package ci

import (
	"list"

	"cue.dev/x/githubactions"
)

workflows: test: githubactions.#Workflow & {
	name: "Test"
	on: {
		push: branches: ["main"]
		pull_request: branches: ["main"]
	}

	defaults: run: shell: "bash"

	jobs: "generate-check": {
		"runs-on": "ubuntu-latest"
		steps: list.Concat([#goSetup, [
			{
				name: "Check the generated workflows"
				run: """
					go generate ./internal/ci
					git diff --exit-code -- .github/workflows
					"""
			},
		]])
	}

	for _, runType in ["ubuntu-latest", "windows-latest"] {
		jobs: "test-\(runType)": {
			"runs-on": runType
			if runType == "windows-latest" {
				// The Windows support is being refined: its failure does not fail the workflow yet.
				"continue-on-error": true
			}
			steps: list.Concat([#goSetup, [
				{
					name: "Run tests"
					env: CHEMONKEY: "1"
					run: "go test ./..."
				},
			]])
		}
	}
}
