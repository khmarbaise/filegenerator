package cmd

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

//CmdText offers the text command with all it's options.
var CmdText = cli.Command{
	Name:        "text",
	Aliases:     []string{"txt"},
	Usage:       "Create a text based file.",
	Description: "text ...",
	Action:      text,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "size",
			Aliases: []string{"s"},
			Usage:   "The size of the file which will be generate",
		},
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "The name of the file.",
		},
	},
}

func text(ctx *cli.Context) error {
	name := "generated.txt"
	if ctx.IsSet("name") {
		name = ctx.String("name")
	}

	size := 200
	if ctx.IsSet("size") {
		size = ctx.Int("size")
	}

	fmt.Printf("The size: %v\n", size)

	file, err := os.Create(fmt.Sprintf("./%s", name))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriterSize(file, 10)
	linesToWrite := []string{"This is an example", "to show how", "to write to a file", "line by line."}
	for _, line := range linesToWrite {
		bytesWritten, err := writer.WriteString(line + "\n")
		if err != nil {
			log.Fatalf("Got error while writing to a file. Err: %s", err.Error())
		}
		fmt.Printf("Bytes Written: %d\n", bytesWritten)
		fmt.Printf("Available: %d\n", writer.Available())
		fmt.Printf("Buffered : %d\n", writer.Buffered())
	}
	writer.Flush()
	return nil
}
