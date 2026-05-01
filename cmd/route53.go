package cmd

import (
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/awsgo/cmd/route53"
)

var Route53Cmd = &cli.Command{
	Name: "route53",
	Usage: `Amazon Route 53 is a highly available and scalable Domain Name System 
(DNS) web service.`,
	EnableShellCompletion: true,
	Commands: []*cli.Command{
		route53.ListHostedZonesCmd,
	},
}
