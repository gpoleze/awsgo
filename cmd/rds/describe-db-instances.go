package rds

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/samber/lo"
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/awsgo/utils"
)

var DescribeDbInstancesCmd = &cli.Command{
	Name:                  "describe-db-instances",
	Aliases:               []string{"list-db-instances", "list-instances", "describe-instances"},
	Category:              "rds",
	EnableShellCompletion: true,
	Flags: []cli.Flag{
		utils.RegionFlag,
		utils.ProfileFlag,
		utils.OutputFlag,
		utils.SortFlag,
	},
	Action: func(ctx context.Context, command *cli.Command) error {
		return utils.WithOutput(ctx, command, DescribeDbInstances, describeDbInstancesToTableRow)
	},
}

func describeDbInstancesToTableRow(instance DBInstance) table.Row {
	return table.Row{
		instance.DBInstanceIdentifier,
		instance.DBInstanceStatus,
		instance.VpcId,
		instance.DBInstanceClass,
		instance.Engine,
		instance.EngineVersion,
		instance.EndpointAddress,
		instance.DBName,
	}
}

func DescribeDbInstances(ctx context.Context, command *cli.Command) ([]DBInstance, error) {
	client, err := utils.GetClient(ctx, command, rds.NewFromConfig)
	if err != nil {
		return nil, err
	}
	output, err := client.DescribeDBInstances(context.TODO(), &rds.DescribeDBInstancesInput{})
	if err != nil {
		fmt.Printf("Couldn't list DB instances: %v\n", err)
		return nil, err
	}

	instances := lo.Map(output.DBInstances, func(item types.DBInstance, _ int) DBInstance {
		return NewDBInstance(item)
	})

	return instances, nil

}
