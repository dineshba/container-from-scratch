package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// go run container.go run <cmd> <args>
// docker run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("invalid command!!")
	}
}
func run() {
	fmt.Printf("Running %v as PID %d \n", os.Args[2:], os.Getpid())
	args := append([]string{"child"}, os.Args[2:]...)
	cmd := exec.Command("/proc/self/exe", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// UTS - Unix Timesharing System
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}
func child() {
	fmt.Printf("Running %v as PID %d \n", os.Args[2:], os.Getpid())
	syscall.Sethostname([]byte("container-demo"))
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	syscall.Chroot("./bundle/rootfs")
	os.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running the command - %s\n", err)
		os.Exit(1)
	}
}
