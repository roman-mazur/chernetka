package ci

// goVersion is the Go release the workflows build with, matching go.mod.
goVersion: "1.27"

// #goSetup checks out the repository and installs Go: the first steps of the jobs building the code.
#goSetup: [...#Step] & [
	{
		name: "Checkout code"
		uses: "actions/checkout@v7"
	},
	{
		name: "Set up Go"
		uses: "actions/setup-go@v7"
		with: "go-version": goVersion
	},
]
