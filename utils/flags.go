package utils

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v3"
	"slices"
)

var RegionFlag = &cli.StringFlag{
	Name:    "region",
	Usage:   "The region to use. Overrides config/env settings.",
	Aliases: []string{"r"},
}

var ProfileFlag = &cli.StringFlag{
	Name:    "profile",
	Usage:   "Use a specific profile from your credential file.",
	Aliases: []string{"p"},
}

var OutputFlag = &cli.StringFlag{
	Name:    "output",
	Usage:   "The formatting style for command output.",
	Value:   "table",
	Aliases: []string{"o"},
	Validator: func(s string) error {
		outoutList := []string{"table", "json"}
		if slices.Contains(outoutList, s) {
			return nil
		}
		return errors.New(fmt.Sprintf("output must be one of %v", outoutList))
	},
}
