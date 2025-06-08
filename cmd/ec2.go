package cmd

import (
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/awsgo/cmd/ec2"
)

var Ec2Cmd = &cli.Command{
	Name:                  "ec2",
	Usage:                 "Amazon Elastic Compute Cloud (Amazon EC2)",
	EnableShellCompletion: true,
	Description: `You can access the features of Amazon Elastic Compute Cloud (Amazon EC2)
		programmatically. For more information, see the Amazon EC2
		Developer Guide .`,
	Commands: []*cli.Command{
		ec2.DescribeVpcsCmd,
		ec2.DescribeInstancesCmd,
	},
}
