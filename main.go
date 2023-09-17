package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker run <container> cmd args
// go run main.go run cmd args
func main() {
	if os.Args[1] == "run" {
		run()
	} else {
		panic(fmt.Sprintf("what? %s", os.Args[1]))
	}
}

func run() {
	fmt.Printf("Running %v as PID %d \n", os.Args[2:], os.Getpid())
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// UTS - Unix Timesharing System
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running the command - %s\n", err)
		os.Exit(1)
	}
}
