package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"math/rand"
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

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func text(ctx *cli.Context) error {
	name := "generated.txt"
	if ctx.IsSet("name") {
		name = ctx.String("name")
	}

	stringLength := 20
	nrOfEntriesPerLine := 10
	lines := 10000

	size := 200
	if ctx.IsSet("size") {
		size = ctx.Int("size")
	}

	fmt.Printf("The size: %v %s\n", size, name)

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
