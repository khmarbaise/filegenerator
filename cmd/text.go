package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"math/rand"
	"os"
)

const (
	optionNumberOfLines          = "lines"
	optionNumberOfEntriesPerLine = "entries"
	optionName                   = "name"
)

//CmdText offers the text command with all it's options.
var CmdText = cli.Command{
	Name:        "text",
	Aliases:     []string{"txt"},
	Usage:       "Create a text based file.",
	Description: "text ...",
	Action:      actionText,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    optionNumberOfLines,
			Aliases: []string{"nol"},
			Usage:   "The number of lines which should be generated",
			Value:   10000,
		},
		&cli.IntFlag{
			Name:        optionNumberOfEntriesPerLine,
			Aliases:     []string{"noe"},
			Usage:       "The number of entries per line",
			Value:       10,
			DefaultText: "10",
		},
		&cli.StringFlag{
			Name:    optionName,
			Aliases: []string{"n"},
			Usage:   "The name of the file.",
		},
	},
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func actionText(ctx *cli.Context) error {
	name := "generated.txt"

	if ctx.IsSet(optionName) {
		name = ctx.String(optionName)
	}

	stringLength := 20

	nrOfEntriesPerLine := 10
	if ctx.IsSet(optionNumberOfEntriesPerLine) {
		nrOfEntriesPerLine = ctx.Int(optionNumberOfEntriesPerLine)
	}

	lines := 10000
	if ctx.IsSet(optionNumberOfLines) {
		lines = ctx.Int(optionNumberOfLines)
	}

	//	fmt.Printf("The size: %v %s\n", size, name)

	file, err := os.Create(fmt.Sprintf("./%s", name))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for lnr := 0; lnr < lines; lnr++ {
		file.WriteString(fmt.Sprintf("%d ", lnr))
		for entry := 0; entry < nrOfEntriesPerLine; entry++ {
			file.WriteString(fmt.Sprintf("%s ", randomString(stringLength)))
		}
		file.WriteString(fmt.Sprintf("\n"))
	}

	return nil
}
