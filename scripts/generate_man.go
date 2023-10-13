package main

import (
	"fmt"

	"github.com/mvisonneau/go-ebsnvme/internal/cli"
)

func main() {
	fmt.Println(cli.NewApp("").ToMan())
}
