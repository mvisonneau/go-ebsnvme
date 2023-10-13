package flags

import (
	"github.com/urfave/cli/v2"

	"github.com/mvisonneau/go-ebsnvme/internal/cli/output"
)

var (
	OutputType = &cli.StringFlag{
		Name:    "output-type",
		Aliases: []string{"t"},
		Usage:   "print results in whether \"text\" or \"json\"",
		Value:   output.TypeText.String(),
	}

	OutputField = &cli.StringSliceFlag{
		Name:    "output-field",
		Aliases: []string{"f"},
		Usage:   "filter out printed fields",
		Value: cli.NewStringSlice(
			output.FieldDeviceName.String(),
			output.FieldDevicePath.String(),
			output.FieldVolumeID.String(),
		),
	}
)
