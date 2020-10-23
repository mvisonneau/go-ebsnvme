package main

import (
	"os"

	"github.com/mvisonneau/go-ebsnvme/internal/cli"
)

var version = ""

func main() {
	cli.Run(version, os.Args)
}
