package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/my-commands/awsgo/utils"
)

var DescribeInstancesCmd = &cli.Command{
	Name:                  "describe-instances",
	Usage:                 "Describes the specified instances or all instances.",
	EnableShellCompletion: true,
	Description: `Describes the specified instances or all instances.

       If you specify instance IDs, the output includes information for only
       the specified instances. If you specify filters, the output includes
       information for only those instances that meet the filter criteria. If
       you do not specify instance IDs or filters, the output includes
       information for all instances, which can affect performance. We
       recommend that you use pagination to ensure that the operation returns
       quickly and successfully.

       If you specify an instance ID that is not valid, an error is returned.
       If you specify an instance that you do not own, it is not included in
       the output.

       Recently terminated instances might appear in the returned results.
       This interval is usually less than one hour.

       If you describe instances in the rare case where an Availability Zone
       is experiencing a service disruption and you specify instance IDs that
       are in the affected zone, or do not specify any instance IDs at all,
       the call fails. If you describe instances and specify only instance IDs
       that are in an unaffected zone, the call works normally.

       The Amazon EC2 API follows an eventual consistency model. This means
       that the result of an API command you run that creates or modifies
       resources might not be immediately available to all subsequent commands
       you run. For guidance on how to manage eventual consistency, see
       Eventual consistency in the Amazon EC2 API in the Amazon EC2 Developer
       Guide. 
		`,
	Flags: []cli.Flag{
		utils.RegionFlag,
		utils.ProfileFlag,
		utils.OutputFlag,
	},
	Aliases:  []string{"list-vpcs"},
	Category: "ec2",
	Action: func(ctx context.Context, command *cli.Command) error {
		return utils.WithOutput[MyInstanceInfo](ctx, command, describeInstances, ec2ItemToTableRow)
	},
}

func describeInstances(ctx context.Context, command *cli.Command) ([]MyInstanceInfo, error) {

	client, errClient := utils.GetClient(ctx, command, ec2.NewFromConfig)
	if errClient != nil {
		return nil, errClient
	}

	input := &ec2.DescribeInstancesInput{}

	result, err := client.DescribeInstances(ctx, input)
	if err != nil {
		return nil, err
	}

	var myInstances []MyInstanceInfo

	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			var myInstance = MyInstanceInfo{
				Name:       utils.FilterTagByKey(i.Tags, "Name"),
				Id:         *i.InstanceId,
				Type:       string(i.InstanceType),
				State:      string(i.State.Name),
				Ami:        *i.ImageId,
				LaunchTime: *i.LaunchTime,
			}
			if i.PrivateIpAddress != nil {
				myInstance.PrivateIp = *i.PrivateIpAddress
			}
			if i.PublicIpAddress != nil {
				myInstance.PublicIp = *i.PublicIpAddress
			}
			myInstances = append(myInstances, myInstance)
		}
	}

	return myInstances, nil
}

//func GetClient(ctx context.Context, command *cli.Command, fromConfig func(cfg aws.Config, optFns ...func(*ec2.Options)) *ec2.Client) (*ec2.Client, error) {
//	region := command.String("region")
//	profile := command.String("profile")
//	cfg, errCfg := config.LoadDefaultConfig(ctx, config.WithRegion(region), config.WithSharedConfigProfile(profile))
//	if errCfg != nil {
//		return nil, errCfg
//	}
//
//	client := fromConfig(cfg)
//	return client, nil
//}

func ec2ItemToTableRow(instance MyInstanceInfo) table.Row {
	return table.Row{
		instance.Name,
		instance.Id,
		instance.Type,
		instance.State,
		instance.Ami,
		instance.LaunchTime,
		instance.PrivateIp,
		instance.PublicIp,
	}
}
