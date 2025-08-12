package main

import (
    "os"
    "context"
	"log"
    "github.com/urfave/cli/v3"
)

func main() {
    if err := (&cli.Command{}).Run(context.Background(), os.Args); err != nil {
    		log.Fatal(err)
	}
}
