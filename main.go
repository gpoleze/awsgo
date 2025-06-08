package main

import (
	"context"
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/my-commands/awsgo/cmd"
	"log"
	"os"
)

var AwsCommands = []*cli.Command{
	cmd.Ec2Cmd,
	cmd.IamCmd,
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
