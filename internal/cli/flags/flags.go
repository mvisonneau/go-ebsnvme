package flags

import (
	"fmt"
	"strings"

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
		Usage:   fmt.Sprintf("filter out fields (%s)", strings.Join(output.GetValidFields().StringSlice(), ",")),
		Value: cli.NewStringSlice(
			output.FieldDeviceName.String(),
			output.FieldVolumeID.String(),
		),
	}
)
