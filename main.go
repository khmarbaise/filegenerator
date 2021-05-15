package main

import (
	"fmt"
	"github.com/khmarbaise/filegenerator/cmd"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

// Version holds the current gjh version
var Version = "development"

// Tags holds the build tags used
var Tags = ""

func main() {
	app := cli.NewApp()
	app.Name = "fgen"
	app.Usage = "How to use fgen"
	app.Description = "file generator"
	app.Version = Version + formatBuiltWith(Tags)
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		&cmd.CmdText,
	}
	app.EnableBashCompletion = true
	err := app.Run(os.Args)
	if err != nil {
		// app.Run already exits for errors implementing ErrorCoder,
		// so we only handle generic errors with code 1 here.
		fmt.Fprintf(app.ErrWriter, "Error: %v\n", err)
		os.Exit(1)
	}

}

func formatBuiltWith(Tags string) string {
	if len(Tags) == 0 {
		return ""
	}

	return " built with: " + strings.Replace(Tags, " ", ", ", -1)
}
