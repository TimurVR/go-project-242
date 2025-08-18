package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"code"
	"github.com/urfave/cli/v3"
)

func main() {
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "help",
	}

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
			if cmd.Bool("help") || cmd.Bool("haaaaalp") || cmd.Bool("halp") {
				(&cli.Command{}).Run(context.Background(), os.Args)
				return nil
			}

			if cmd.NArg() == 0 {
				return fmt.Errorf("path argument is required")
			}

			path := cmd.Args().First()
			humanReadable := cmd.Bool("human")
			allReadable := cmd.Bool("all")
			recursionReadable := cmd.Bool("recursive")
			size, err := code.GetPathSize(path, recursionReadable, humanReadable, allReadable)
			if err != nil {
				return fmt.Errorf("error getting size: %w", err)
			}			
			fmt.Printf("%s\t%s\n", size, path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
