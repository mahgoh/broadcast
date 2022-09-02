package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"time"

	"github.com/mahgoh/bm-ws/broadcasts/internal/broadcast"
	"github.com/morikuni/aec"
)

type Command struct {
	usage    string
	doneMsg  string
	showDone bool
	Run      func([]string)
}

func (c *Command) Done(d time.Duration) {
	if c.showDone {
		fmt.Printf("[%sDONE%s] %s %dms\n", aec.GreenF, aec.Reset, c.doneMsg, d.Milliseconds())
	}
}

func createCommands() map[string]*Command {
	commands := make(map[string]*Command)

	commands["create"] = cmdCreate()
	commands["v"] = cmdVersion()
	commands["version"] = cmdVersion()

	return commands
}

// cmdCreate creates a broadcast
func cmdCreate() *Command {
	return &Command{
		usage:    "create <path>",
		doneMsg:  "Create broadcast.",
		showDone: true,
		Run: func(args []string) {
			if len(args) < 2 {
				fmt.Println("No path specified.")
				return
			}

			b := broadcast.NewBroadcast(args[1])
			b.Transform()
			buf := b.Parse()

			filePath := path.Join(args[1], "broadcast.html")
			if err := ioutil.WriteFile(filePath, buf.Bytes(), 0777); err != nil {
				log.Fatal(err)
			}

		},
	}
}

// cmdVersion prints the version
func cmdVersion() *Command {
	return &Command{
		usage:    "v version",
		showDone: false,
		Run: func(_ []string) {
			fmt.Println(Version)
		},
	}
}
