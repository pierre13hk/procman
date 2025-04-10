package main

import (
    "testing"
)

func TestProcmanConfFromJsonOk(t *testing.T) {
    json := `
    {
        "working_dir": "/tmp",
        "procman_dir": "/procman",
        "process_definitions": [
            {
                "name": "test_process",
                "cmd": "echo",
                "args": ["Hello, World!"],
                "env": {
                    "TEST_ENV": "test_value"
                }
            }
        ]
    }` // Properly close the JSON string

    config, err := procmanConfFromJson(json)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    // Add assertions to validate the parsed config
    if config.WorkingDir != "/tmp" {
        t.Errorf("Expected working_dir to be /tmp, got %s", config.WorkingDir)
    }
    if config.ProcmanDir != "/procman" {
        t.Errorf("Expected procman_dir to be /procman, got %s", config.ProcmanDir)
    }
}

func TestProcmanConfFromJsonMissingField(t *testing.T) {
	json := `
	{
		"working_dir": "/tmp",
		"procman_dir": "/procman"
	}` // Missing process_definitions

	_, err := procmanConfFromJson(json)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}