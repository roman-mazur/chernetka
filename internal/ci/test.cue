package ci

import "list"

workflows: test: #Workflow & {
	name: "Test"
	on: {
		push: branches: ["main"]
		pull_request: branches: ["main"]
	}
	jobs: test: {
		"runs-on": "ubuntu-latest"
		steps: list.Concat([#goSetup, [
			{
				name: "Run tests"
				run:  "CHEMONKEY=1 go test ./..."
			},
			{
				name: "Check the generated workflows"
				run: """
					go generate ./internal/ci
					git diff --exit-code -- .github/workflows
					"""
			},
		]])
	}
}
