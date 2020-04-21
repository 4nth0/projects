package main

import "flag"

func addCmd() command {
	fs := flag.NewFlagSet("projects add", flag.ExitOnError)

	return command{fs, func(args []string) error {
		return addProject()
	}}
}

func addProject() error {

	// project add project
	// project add tag
	// project rm tag (list and select tag to delete)
	// project set description

	// 1. Load Configuration
	// 2. Use the current directory ?
	// 3.

	return nil
}
