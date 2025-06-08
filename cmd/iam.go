package cmd

import (
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/my-commands/awsgo/cmd/iam"
)

var IamCmd = &cli.Command{
	Name:                  "iam",
	Usage:                 "Identity and Access Management (IAM)",
	EnableShellCompletion: true,
	Description: `Identity and Access Management (IAM) is a web service for securely
       controlling access to Amazon Web Services services. With IAM, you can
       centrally manage users, security credentials such as access keys, and
       permissions that control which Amazon Web Services resources users and
       applications can access. For more information about IAM, see Identity
       and Access Management (IAM) and the Identity and Access Management User
       Guide .`,
	Commands: []*cli.Command{
		iam.GetGroupCmd,
	},
}
