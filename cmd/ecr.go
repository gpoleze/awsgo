package cmd

import (
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/my-commands/awsgo/cmd/ecr"
)

var EcrCmd = &cli.Command{
	Name:                  "ecr",
	Usage:                 "Amazon  Elastic  Container Registry (Amazon ECR)",
	EnableShellCompletion: true,
	Commands: []*cli.Command{
		ecr.DescribeRepositoriesCmd,
	},
}
