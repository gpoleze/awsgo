package ecr

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/samber/lo"
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/my-commands/awsgo/utils"
)

var DescribeRepositoriesCmd = &cli.Command{
	Name:                  "describe-repositories",
	Aliases:               []string{"list-repositories", "list-repos", "describe-repos"},
	EnableShellCompletion: true,
	Flags: []cli.Flag{
		utils.RegionFlag,
		utils.ProfileFlag,
		utils.OutputFlag,
		utils.SortFlag,
	},
	Category: "ecr",
	Action: func(ctx context.Context, command *cli.Command) error {
		_ = utils.SortFlag.Set("Value", "REPOSITORYNAME")
		return utils.WithOutput(ctx, command, describeRepositories, describeRepositoriesItemToTableRow)
	},
}

func describeRepositories(ctx context.Context, command *cli.Command) ([]Repository, error) {
	client, err := utils.GetClient(ctx, command, ecr.NewFromConfig)
	if err != nil {
		return nil, err
	}

	input := &ecr.DescribeRepositoriesInput{}

	result, err := client.DescribeRepositories(ctx, input)
	if err != nil {
		fmt.Printf("failed to describe repositories, %v\n", err)
		return nil, err
	}
	repositories := lo.Map(result.Repositories, func(item types.Repository, _ int) Repository { return NewRepository(item) })

	return repositories, nil
}

func describeRepositoriesItemToTableRow(repository Repository) table.Row {
	return table.Row{
		repository.RepositoryName,
		repository.RepositoryUri,
		repository.CreatedAt,
	}
}
