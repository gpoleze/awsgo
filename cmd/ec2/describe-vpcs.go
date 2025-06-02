package ec2

import (
	"context"
	"sort"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v3"

	"gitlab.com/gabriel.poleze/my-commands/awsgo/utils"
)

var DescribeVpcsCmd = &cli.Command{
	Name:    "describe-vpcs",
	Aliases: []string{"list-vpcs"},
	Usage:   "Describe your VPCs",
	Description: `
		Describes your VPCs. The default is to describe all your VPCs.
		Alternatively, you can specify specific VPC IDs or filter the results
		to include only the VPCs that match specific criteria.

		See also: AWS API Documentation
		`,
	Category:              "ec2",
	EnableShellCompletion: true,
	Flags: []cli.Flag{
		utils.RegionFlag,
		utils.ProfileFlag,
		utils.OutputFlag,
	},
	Action: func(ctx context.Context, command *cli.Command) error {
		return utils.WithOutput(ctx, command, describeVpcs, vpcItemToTableRow)
	},
}

func describeVpcs(ctx context.Context, cmd *cli.Command) ([]Vpc, error) {
	client, err := utils.GetClient(ctx, cmd, ec2.NewFromConfig)
	if err != nil {
		return nil, err
	}

	input := &ec2.DescribeVpcsInput{}
	result, err := client.DescribeVpcs(context.TODO(), input)
	if err != nil {
		return nil, err
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

	return vpcs, nil
}

func vpcItemToTableRow(vpc Vpc) table.Row {
	return table.Row{
		vpc.Name,
		vpc.Id,
		vpc.CidrBlock,
	}
}
