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
}

func main() {
	cmd := &cli.Command{
		Name:                  "awsgo",
		Usage:                 "AWS cli with a more user friendly output",
		Commands:              AwsCommands,
		EnableShellCompletion: true,
		Suggest:               true,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
