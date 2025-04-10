package main

import (
)

type ProcessDefinition struct {
	// The name of the process
	Name string `json:"name"`
	// The command to run the process
	Cmd string `json:"cmd"`
	// The arguments to pass to the command
	Args []string `json:"args"`
	// The environment variables to set for the process
	Env map[string]string `json:"env"`
}
