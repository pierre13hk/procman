package main

import (
	"log"
	"os"
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

type ProcmanProcess struct {
	osProc     *os.Process
	definition *ProcessDefinition
}

func (proc *ProcmanProcess) start() error {
	// Implement the logic to start the process
	// This is a placeholder implementation
	strEnv := make([]string, 0, len(proc.definition.Env))
	for k, v := range proc.definition.Env {
		strEnv = append(strEnv, k+"="+v)
	}
	osProc, err := os.StartProcess(proc.definition.Cmd, proc.definition.Args, &os.ProcAttr{
		Env: strEnv,
		Files: []*os.File{
			os.Stdin,  // Standard input
			os.Stdout, // Standard output
			os.Stderr, // Standard error
		},
	})
	if err != nil {
		log.Printf("Error starting process %s: %v", proc.definition.Name, err)
		return err
	}
	proc.osProc = osProc
	log.Printf("Started process %s with PID %d", proc.definition.Name, osProc.Pid)
	return nil
}

func (proc *ProcmanProcess) wait() int {
	// Implement the logic to wait for the process to finish
	// This is a placeholder implementation
	if proc == nil {
		log.Println("Process is nil, cannot wait")
		return -1
	}
	state, err := proc.osProc.Wait()
	pid := proc.osProc.Pid
	proc.osProc = nil
	if err != nil {
		log.Printf("Error waiting for process %d: %v", pid, err)
		// creates dangling process?
		return -1
	}
	log.Printf("Process %d finished with state: %v", pid, state)
	return state.ExitCode()
}

func (proc *ProcmanProcess) manage() {
	// Implement the logic to manage the process
	if proc == nil {
		log.Println("Process is nil, cannot manage")
		return
	}
	log.Printf("Managing process %s", proc.definition.Name)
	for {
		if proc.osProc == nil {
			proc.start()
		}
		exitCode := proc.wait()
		if exitCode != 0 {
			log.Printf("Process %s exited with code %d, restarting...", proc.definition.Name, exitCode)
		} else {
			log.Printf("Process %s exited normally, not restarting", proc.definition.Name)
			break
		}
	}
}
