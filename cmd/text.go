package cmd

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	optionNumberOfLines          = "lines"
	optionNumberOfEntriesPerLine = "entries"
	optionName                   = "name"
	optionSize                   = "size"
	buffer64k                    = 64 * 1024
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
			Name:    optionNumberOfEntriesPerLine,
			Aliases: []string{"noe"},
			Usage:   "The number of entries per line",
			Value:   10,
		},
		&cli.IntFlag{
			Name:    optionSize,
			Aliases: []string{"s"},
			Usage:   "The number of characters per string",
			Value:   20,
		},
		&cli.StringFlag{
			Name:    optionName,
			Aliases: []string{"n"},
			Usage:   "The name of the file.",
			Value:   "generated.txt",
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

	rand.Seed(time.Now().UnixNano())

	name := ctx.String(optionName)

	stringLength := ctx.Int(optionSize)

	nrOfEntriesPerLine := ctx.Int(optionNumberOfEntriesPerLine)

	lines := ctx.Int(optionNumberOfLines)

	fmt.Printf("Generating file %s with %d lines containing %d entries per line which comprises of %d characters\n", name, lines, nrOfEntriesPerLine, stringLength)

	file, err := os.Create(fmt.Sprintf("./%s", name))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriterSize(file, buffer64k)

	defer bufferedWriter.Flush()

	for lnr := 0; lnr < lines; lnr++ {
		bufferedWriter.WriteString(fmt.Sprintf("%d ", lnr))
		for entry := 0; entry < nrOfEntriesPerLine; entry++ {
			bufferedWriter.WriteString(fmt.Sprintf("%s ", randomString(stringLength)))
		}
		bufferedWriter.WriteString(fmt.Sprintf("\n"))
	}

	return nil
}
