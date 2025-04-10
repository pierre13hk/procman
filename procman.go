package main 

import (
	"encoding/json"
	"fmt"
	"os"
)

type ProcmanConfig struct {
	// The process definitions
	ProcessDefinitions []ProcessDefinition `json:"process_definitions"`
	// The process manager's working directory
	WorkingDir string `json:"working_dir"`
	// The process manager's directory
	ProcmanDir string `json:"procman_dir"`
}

func (conf *ProcmanConfig) validate() error {
	// Validate the configuration
	// Check if the working directory is set
	if conf.WorkingDir == "" {
		return fmt.Errorf("working_dir is not set")
	}
	// Check if the process manager directory is set
	if conf.ProcmanDir == "" {
		return fmt.Errorf("procman_dir is not set")
	}
	// Check if the process definitions are set
	if len(conf.ProcessDefinitions) == 0 {
		return fmt.Errorf("process_definitions is not set")
	}
	return nil
}

func procmanConfFromJsonFile(filepath string) (ProcmanConfig, error) {
	// Read the JSON file and unmarshal it into a ProcmanConfig struct
	// Return the ProcmanConfig struct and any error that occurred
	json_contents, error := os.ReadFile(filepath)
	if error != nil {
		return ProcmanConfig{}, error
	}
	return procmanConfFromJson(string(json_contents))
}

func procmanConfFromJson(strConf string) (ProcmanConfig, error) {
	// Unmarshal the JSON string into a ProcmanConfig struct
	// Return the ProcmanConfig struct and any error that occurred
	var config ProcmanConfig
	err := json.Unmarshal([]byte(strConf), &config)
	if err != nil {
		return ProcmanConfig{}, err
	}

	if err = config.validate(); err != nil {
		return ProcmanConfig{}, err
	}

	return config, nil
}
