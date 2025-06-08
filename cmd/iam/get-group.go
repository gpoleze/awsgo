package iam

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/samber/lo"
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/my-commands/awsgo/utils"
)

var groupNameFlag = &cli.StringFlag{
	Name: "group-name",
	Usage: `The name of the group.

          This parameter allows (through its regex pattern ) a string of
          characters consisting of upper and lowercase alphanumeric characters
          with no spaces. You can also include any of the following
          characters: _+=,.@-`,
	Aliases:  []string{"g"},
	Required: true,
}

var GetGroupCmd = &cli.Command{
	Name:                  "get-group",
	Usage:                 "Get information about the specified IAM group.",
	Description:           `Returns a list of IAM users that are in the specified IAM group.`,
	Category:              "iam",
	EnableShellCompletion: true,
	Flags: []cli.Flag{
		groupNameFlag,
		utils.RegionFlag,
		utils.ProfileFlag,
		utils.OutputFlag,
		utils.SortFlag,
	},
	Action: func(ctx context.Context, command *cli.Command) error {
		_ = utils.SortFlag.Set("Value", "username")
		return utils.WithOutput(ctx, command, GetGroup, getGroupItemToTableRow)
	},
}

func getGroupItemToTableRow(user IamUser) table.Row {
	return table.Row{
		user.UserName,
		user.UserId,
		user.Arn,
		user.CreateDate,
		user.PasswordLastUsed,
	}
}

func GetGroup(ctx context.Context, command *cli.Command) ([]IamUser, error) {
	client, err := utils.GetClient(ctx, command, iam.NewFromConfig)
	if err != nil {
		return nil, err
	}
	groupName := command.String(groupNameFlag.Name)
	input := &iam.GetGroupInput{
		GroupName: &groupName,
	}

	result, err := client.GetGroup(ctx, input)
	if err != nil {
		fmt.Println("failed to get group users, %v", err)
		return nil, err
	}

	users := lo.Map(result.Users, func(item types.User, index int) IamUser {
		return NewIamUser(item)
	})

	return users, nil
}
