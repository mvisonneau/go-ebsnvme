package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/mvisonneau/go-ebsnvme/pkg/ebsnvme"
	"github.com/urfave/cli"
)

const (
	usage = "go-ebsnvme <block_device> [--volume-id|--device-name]"
)

// Run handles the instanciation of the CLI application
func Run(version string) {
	NewApp(version, time.Now()).Run(os.Args)
}

// NewApp configures the CLI application
func NewApp(version string, start time.Time) (app *cli.App) {
	app = cli.NewApp()
	app.Name = "go-ebsnvme"
	app.Version = version
	app.Usage = "Fetch information about AWS EBS NVMe volumes"
	app.UsageText = usage
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "volume-id, i",
			Usage: "only print the EBS volume-id",
		},
		cli.BoolFlag{
			Name:  "device-name, n",
			Usage: "only print the name of the block device",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		if len(ctx.Args()) != 1 ||
			(ctx.Bool("volume-id") && ctx.Bool("device-name")) {
			return cli.NewExitError("Usage: "+usage, 1)
		}

		d, err := ebsnvme.ScanDevice(ctx.Args().First())
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return cli.NewExitError("", 1)
		}

		if ctx.Bool("volume-id") {
			fmt.Println(d.VolumeID)
			return nil
		}

		if ctx.Bool("device-name") {
			fmt.Println(d.Name)
			return nil
		}

		fmt.Println(d.VolumeID)
		fmt.Println(d.Name)
		return nil
	}

	app.Metadata = map[string]interface{}{
		"startTime": start,
	}

	return
}
