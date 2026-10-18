package ci

// A minimal schema of the GitHub Actions workflows: the definitions are closed,
// so a misspelled field fails the generation.
// TODO: switch to cue.dev/x/githubactions.

#Workflow: {
	name: string
	on: {...}
	permissions?: #Permissions
	env?: [string]: string
	jobs: [string]: #Job
}

#Job: {
	"runs-on": string
	needs?: string | [...string]
	if?:          string
	permissions?: #Permissions
	env?: [string]:     string
	outputs?: [string]: string
	steps: [...#Step]
}

#Step: {
	name:  string
	id?:   string
	if?:   string
	uses?: string
	with?: [string]: string | number | bool
	env?: [string]:  string
	run?: string
}

#Permissions: [string]: "read" | "write" | "none"
