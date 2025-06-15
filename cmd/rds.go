package cmd

import (
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/awsgo/cmd/rds"
)

var RdsCmd = &cli.Command{
	Name:                  "rds",
	Usage:                 "Amazon Relational Database Service (Amazon RDS)",
	EnableShellCompletion: true,
	Commands: []*cli.Command{
		rds.DescribeDbInstancesCmd,
	},
}
