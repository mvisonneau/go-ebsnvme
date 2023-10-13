package cli

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/mvisonneau/go-ebsnvme/internal/cli/flags"
	"github.com/mvisonneau/go-ebsnvme/internal/cli/output"
	"github.com/mvisonneau/go-ebsnvme/pkg/ebsnvme"
)

const (
	usageText = "go-ebsnvme [opts] <block_device>"
)

// Run handles the instantiation of the CLI application.
func Run(version string, args []string) {
	err := NewApp(version).Run(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// NewApp configures the CLI application.
func NewApp(version string) (app *cli.App) {
	app = cli.NewApp()
	app.Name = "go-ebsnvme"
	app.Version = version
	app.Usage = "Find details about currently attached AWS EBS NVMe volumes"
	app.UsageText = usageText
	app.EnableBashCompletion = true

	app.Flags = cli.FlagsByName{
		flags.OutputType,
		flags.OutputField,
	}

	app.Action = func(ctx *cli.Context) (err error) {
		var outputType output.Type
		if outputType, err = output.ParseTypeFromString(flags.OutputType.Get(ctx)); err != nil {
			return
		}

		var outputFields output.Fields
		if outputFields, err = output.ParseFieldsFromStringSlice(flags.OutputField.Get(ctx)); err != nil {
			return
		}

		var device ebsnvme.Device
		if device, err = ebsnvme.ScanDevice(ctx.Args().First()); err != nil {
			return
		}

		fmt.Println(output.FormatDeviceDetails(
			device,
			outputType,
			outputFields,
		))

		return
	}

	return
}
