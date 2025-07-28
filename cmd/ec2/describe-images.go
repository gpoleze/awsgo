package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/samber/lo"
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/awsgo/utils"
	"time"
)

var sinceFlag = &cli.TimestampFlag{
	Name: "since",
	Config: cli.TimestampConfig{
		Layouts: []string{"2006-01-02T15:04:05"},
	},
	Aliases: []string{"after"},
	Value:   time.Now().Add(-30),
}

var DescribeImagesCmd = &cli.Command{
	Name:                  "describe-images",
	Usage:                 "Describes all private images (AMIs, AKIs, and ARIs) available to you.",
	EnableShellCompletion: true,
	Flags: []cli.Flag{
		utils.RegionFlag,
		utils.ProfileFlag,
		utils.OutputFlag,
		utils.SortFlag,
		sinceFlag,
	},

	Aliases:  []string{"list-images"},
	Category: "ec2",
	Action: func(ctx context.Context, command *cli.Command) error {
		return utils.WithTableOutput[Image](ctx, command, describeImages, utils.BuildTableParams[Image]{
			Title:             "Ami Instances",
			ItemToRowFunction: ec2ImageToTableRow,
			SortBy: []table.SortBy{table.SortBy{
				Name:       "creation_date",
				IgnoreCase: true,
				Mode:       table.Asc,
			}},
		})
	},
}

func describeImages(ctx context.Context, command *cli.Command) ([]Image, error) {

	client, errClient := utils.GetClient(ctx, command, ec2.NewFromConfig)
	if errClient != nil {
		return nil, errClient
	}

	imageIsPublicFilterName := "is-public"
	imageIsPublicFilter := types.Filter{Name: &imageIsPublicFilterName, Values: []string{"false"}}

	result, err := client.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Filters: []types.Filter{
			imageIsPublicFilter,
		},
	})
	if err != nil {
		return nil, err
	}

	images := lo.Map(result.Images, func(item types.Image, _ int) Image {
		return NewImage(item)
	})

	return images, nil
}

func ec2ImageToTableRow(item Image) table.Row {
	return table.Row{
		item.Name,
		item.ImageId,
		item.ImageType,
		item.CreationDate,
		item.State,
	}
}
