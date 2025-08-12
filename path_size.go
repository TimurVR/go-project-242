package goproject242

import (
    "fmt"
    "log"
    "os"
    "context"

    "github.com/urfave/cli/v3"
)

func _() {
    cmd := &cli.Command{
        Name:  "help",
        Usage: "make an explosive entrance",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "help",
				Aliases: []string{"h"},
				Usage:   "show help",
			},
		},
        Action: func(context.Context, *cli.Command) error {
			Info()
            return nil
        },
    }

    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}
func Info(){
	fmt.Print(`NAME:
    hexlet-path-size - print size of a file or directory;

USAGE:
    hexlet-path-size [global options]

GLOBAL OPTIONS:
    --help, -h  show help
`)
}
