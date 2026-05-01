package route53

import (
	"context"
	"slices"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/awsgo/utils"
)

var ListHostedZonesCmd = &cli.Command{
	Name: "list-hosted-zones",
	Flags: []cli.Flag{
		utils.ProfileFlag,
		utils.OutputFlag,
		utils.SortFlag,
	},
	Category: "route53",
	Action: func(ctx context.Context, command *cli.Command) error {
		return utils.WithOutput[HostedZone](ctx, command, listHostedZones, hostedZoneToTableRow)
	},
}

func listHostedZones(ctx context.Context, command *cli.Command) ([]HostedZone, error) {
	client, errClient := utils.GetClient(ctx, command, route53.NewFromConfig)
	if errClient != nil {
		return nil, errClient

	}

	result, err := client.ListHostedZones(ctx, &route53.ListHostedZonesInput{})
	if err != nil {
		return nil, err
	}

	hostedZones := make([]HostedZone, len(result.HostedZones))
	for i, hzItem := range result.HostedZones {
		hostedZones[i] = NewHostedZone(hzItem)
	}
	slices.SortFunc(hostedZones, func(a, b HostedZone) int {
		return strings.Compare(a.Name, b.Name)
	})
	return hostedZones, nil
}

func hostedZoneToTableRow(hz HostedZone) table.Row {
	return table.Row{
		hz.Id,
		hz.Name,
		hz.ResourceRecordSetCount,
		hz.Config.PrivateZone,
		hz.Config.Comment,
	}
}
