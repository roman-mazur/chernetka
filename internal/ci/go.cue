package ci

import "cue.dev/x/githubactions"

// #goSetup checks out the repository and installs Go, in the version go.mod declares: the first
// steps of the jobs building the code.
#goSetup: [...githubactions.#Step] & [
	{
		name: "Checkout code"
		uses: "actions/checkout@v7"
	},
	{
		name: "Set up Go"
		uses: "actions/setup-go@v7"
		with: "go-version-file": "go.mod"
	},
]
