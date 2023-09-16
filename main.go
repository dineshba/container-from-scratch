package main

import (
	"fmt"
	"os"
	"os/exec"
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
	cmd.Run()
}
