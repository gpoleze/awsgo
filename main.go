package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/awsgo/cmd"
)

var AwsCommands = []*cli.Command{
	cmd.Ec2Cmd,
	cmd.IamCmd,
	cmd.EcrCmd,
	cmd.RdsCmd,
	cmd.SesCmd,
}

func main() {
	command := &cli.Command{
		Name:                  "awsgo",
		Usage:                 "AWS cli with a more user friendly output",
		Commands:              AwsCommands,
		EnableShellCompletion: true,
		Suggest:               true,
	}

	if err := command.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
