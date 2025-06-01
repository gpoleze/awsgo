package ec2

import (
	"context"
	"sort"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v3"

	"gitlab.com/gabriel.poleze/my-commands/awsgo/utils"
)

var DescribeVpcsCmd = &cli.Command{
	Name: "describe-vpcs",
	Description: `
		Describes your VPCs. The default is to describe all your VPCs.
		Alternatively, you can specify specific VPC IDs or filter the results
		to include only the VPCs that match specific criteria.

		See also: AWS API Documentation

		describe-vpcs is a paginated operation. Multiple API calls may be
		issued in order to retrieve the entire data set of results. You can
		disable pagination by providing the --no-paginate argument.  When using
		--output text and the --query argument on a paginated response, the
		--query argument must extract data from the results of the following
		query expressions: Vpcs
		`,
	Flags: []cli.Flag{
		utils.RegionFlag,
		utils.ProfileFlag,
		utils.OutputFlag,
	},
	Aliases:  []string{"list-vpcs"},
	Category: "ec2",
	Action:   describeVpcs,
}

func describeVpcs(ctx context.Context, cmd *cli.Command) error {
	region := cmd.String("region")
	profile := cmd.String("profile")

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithSharedConfigProfile(profile))
	if err != nil {
		panic(err)
	}
	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeVpcsInput{}
	result, err := client.DescribeVpcs(context.TODO(), input)
	if err != nil {
		return err
	}

	var vpcs = Vpcs{}
	cidrs := []string{}
	for _, i := range result.Vpcs {
		vpc := Vpc{
			Name:      utils.FilterTagByKey(i.Tags, "Name"),
			Id:        *i.VpcId,
			CidrBlock: *i.CidrBlock,
		}
		vpcs = append(vpcs, vpc)
		cidrs = append(cidrs, *i.CidrBlock)
	}

	sort.Sort(vpcs)

	switch cmd.String("output") {
	case "json":
		utils.PrintJson(vpcs)
	case "table":
		utils.BuildTable(vpcs, itemToTableRow)
	default:
		utils.BuildTable(vpcs, itemToTableRow)
	}
	return nil
}

func itemToTableRow(vpc Vpc) table.Row {
	return table.Row{
		vpc.Name,
		vpc.Id,
		vpc.CidrBlock,
	}
}
