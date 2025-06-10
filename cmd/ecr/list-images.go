package ecr

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/samber/lo"
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/awsgo/utils"
)

var groupNameFlag = &cli.StringFlag{
	Name:     "repository-name",
	Usage:    "The repository with image IDs to be listed.",
	Aliases:  []string{"rn"},
	Required: true,
}

var ListImagesCmd = &cli.Command{
	Name:                  "list-images",
	Description:           `Lists all the image IDs for the specified repository`,
	EnableShellCompletion: true,
	Flags: []cli.Flag{
		utils.RegionFlag,
		utils.ProfileFlag,
		utils.OutputFlag,
		utils.SortFlag,
		utils.SortOrderFlag,
		groupNameFlag,
	},
	Category: "ecr",
	Action: func(ctx context.Context, command *cli.Command) error {
		_ = utils.SortFlag.Set("Value", "pushed_at")
		_ = utils.SortOrderFlag.Set("Value", "dsc")
		return utils.WithOutput(ctx, command, listImages, listImagesItemToTableRow)
	},
}

func listImages(ctx context.Context, command *cli.Command) ([]Image, error) {
	client, err := utils.GetClient(ctx, command, ecr.NewFromConfig)
	if err != nil {
		return nil, err
	}

	repositoryName := command.String(groupNameFlag.Name)

	input := &ecr.DescribeImagesInput{
		RepositoryName: &repositoryName,
	}

	result, err := client.DescribeImages(ctx, input)
	if err != nil {
		fmt.Printf("failed to list images, %v\n", err)
		return nil, err
	}

	images := lo.Map(result.ImageDetails, func(item types.ImageDetail, _ int) Image {
		return NewImage(item)
	})

	return images, nil
}

func listImagesItemToTableRow(image Image) table.Row {
	return table.Row{
		image.Tags,
		image.PushedAt,
		image.MediaType,
		image.SizeMB,
		image.Digest,
	}
}
