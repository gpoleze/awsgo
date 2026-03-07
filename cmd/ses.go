package cmd

import (
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/awsgo/cmd/ses"
)

var SesCmd = &cli.Command{
	Name:                  "ses",
	Usage:                 "Amazon  SES  is  an  Amazon Web Services service that you can use to\n          send email messages to your customers.",
	EnableShellCompletion: true,
	Commands: []*cli.Command{
		ses.DescribeInstancesCmd,
	},
}
