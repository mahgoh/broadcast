package main

import (
	"fmt"
	"os"
	"time"
)

var (
	Version string
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No command specified.")
		return
	}

	cmds := createCommands()

	if cmd, ok := cmds[args[0]]; ok {
		start := time.Now()
		cmd.Run(args)
		duration := time.Since(start)
		cmd.Done(duration)
		return
	}

	fmt.Println("Command does not exist.")
}
