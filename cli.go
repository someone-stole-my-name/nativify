package main

import (
	"fmt"
	"os"

	"github.com/someone-stole-my-name/nativify/command"

	"github.com/mitchellh/cli"
)

func Run(args []string) int {

	// Meta-option for executables.
	// It defines output color and its stdout/stderr stream.
	meta := &command.Meta{
		Ui: &cli.ColoredUi{
			InfoColor:  cli.UiColorBlue,
			ErrorColor: cli.UiColorRed,
			Ui: &cli.BasicUi{
				Writer:      os.Stdout,
				ErrorWriter: os.Stderr,
				Reader:      os.Stdin,
			},
		}}

	return RunCustom(args, Commands(meta))
}

func RunCustom(args []string, commands map[string]cli.CommandFactory) int {

	// Get the command line args. We shortcut "--version" and "-v" to
	// just show the version.
	for _, arg := range args {
		if arg == "-v" || arg == "-version" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}
	}

	if len(args) > 1 && args[0] == "add" {
		newArgs := make([]string, 4)
		newArgs[0] = args[0]
		for i, arg := range args {
			if arg == "-n" || arg == "-name" || arg == "--name" {
				newArgs[1] = args[i+1]
			} else if arg == "-u" || arg == "-url" || arg == "--url" {
				newArgs[2] = args[i+1]
			} else if arg == "-i" || arg == "-icon" || arg == "--icon" {
				newArgs[3] = args[i+1]
			}
		}
		args = newArgs
	}

	cli := &cli.CLI{
		Args:       args,
		Commands:   commands,
		Version:    Version,
		HelpFunc:   cli.BasicHelpFunc(Name),
		HelpWriter: os.Stdout,
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute: %s\n", err.Error())
	}

	return exitCode
}
