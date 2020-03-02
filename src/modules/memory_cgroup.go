package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

const memoryHierarchyMount = "/sys/fs/cgroup/memory"

func main() {

	if os.Args[0] == "/proc/self/exe" {
		// this is the code block for forked container process

		fmt.Printf("Current pid = %d", syscall.Getpid())
		fmt.Println()

		cmd := exec.Command("sh", "-c", "stress --vm-bytes 10m --vm-keep -m 1")
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	cmd := exec.Command("/proc/self/exe") // exec a clone of "self"
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	} else {
		fmt.Printf("Container pid = %v", cmd.Process.Pid)
		fmt.Println()

		os.Mkdir(path.Join(memoryHierarchyMount, "test-memory-limit"), 0755)
		ioutil.WriteFile(
			path.Join(memoryHierarchyMount, "test-memory-limit", "tasks"),
			[]byte(strconv.Itoa(cmd.Process.Pid)),
			0644,
		)
		ioutil.WriteFile(
			path.Join(memoryHierarchyMount, "test-memory-limit", "memory.limit_in_bytes"),
			[]byte("100m"),
			0644,
		)
	}
	cmd.Process.Wait()
}
