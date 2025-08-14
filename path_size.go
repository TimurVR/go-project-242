package goproject242

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func Main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "help",
				Aliases: []string{"h"},
				Usage:   "show help",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},				
				Usage:   "include hidden files and directories",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},				
				Usage:   "recursive size of directories",
				Value:   false,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.Bool("help") {
				Info()
				return nil
			}

			if cmd.NArg() == 0 {
				return fmt.Errorf("path argument is required")
			}

			path := cmd.Args().First()
			humanReadable := cmd.Bool("human")
			allReadable := cmd.Bool("all")
			recursionReadable := cmd.Bool("recursive")
			size, err := GetPathSize(path, recursionReadable, humanReadable, allReadable)
			if err != nil {
				return fmt.Errorf("error getting size: %w", err)
			}
			
			fmt.Println(size)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
func Info(){
	fmt.Print(`NAME:
   hexlet-path-size - print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)

USAGE:
   hexlet-path-size [global options]

GLOBAL OPTIONS:
   --recursive, -r  recursive size of directories (default: false)
   --human, -H      human-readable sizes (auto-select unit) (default: false)
   --all, -a        include hidden files and directories (default: false)
   --help, -h       show help
`)
}