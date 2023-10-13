package main

import (
	"fmt"
	"time"

	"github.com/mvisonneau/go-ebsnvme/internal/cli"
)

func main() {
	fmt.Println(cli.NewApp("", time.Time{}).ToMan())
}
