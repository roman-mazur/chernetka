package ci

import (
	"encoding/yaml"
	"tool/file"
)

// gen writes the workflows to .github/workflows: go generate ./internal/ci
command: gen: {
	for name, w in workflows {
		"write-\(name)": file.Create & {
			filename: "../../.github/workflows/\(name).yml"
			contents: """
				# Code generated from internal/ci with go generate; DO NOT EDIT.

				\(yaml.Marshal(w))
				"""
		}
	}
}
