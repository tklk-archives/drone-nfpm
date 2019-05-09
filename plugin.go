package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type (
	Plugin struct {
		Config string
		Target string
	}
)

var nfpmExecutable = "nfpm"

// Exec executes the plugins functionality
func (p Plugin) Exec() error {
	var cmds = make([]*exec.Cmd, 0)

	cmds = append(cmds, commandBuild(p))
	return execAll(cmds)
}

func commandBuild(p Plugin) *exec.Cmd {
	var args = make([]string, 0)

	args = append(args, "pkg")

	if p.Config != "" {
		args = append(args, "--config", p.Config)
	}

	if p.Target != "" {
		args = append(args, "--target", p.Target)
	}

	return exec.Command(nfpmExecutable, args...)
}

// trace writes each command to stdout with the command wrapped in an xml
// tag so that it can be extracted and displayed in the logs.
func trace(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "+ %s\n", strings.Join(cmd.Args, " "))
}

// execAll executes a slice of commands as a batch job
func execAll(cmds []*exec.Cmd) error {
	// Execute all commands
	for _, cmd := range cmds {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		trace(cmd)

		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
