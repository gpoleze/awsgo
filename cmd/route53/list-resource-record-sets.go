package route53

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/awsgo/utils"
)

var HostedZoneIdFlag = &cli.StringFlag{
	Name:    "hosted-zone-id",
	Usage:   "The hosted zone ID to use for hosted zones",
	Aliases: []string{"i"},
}

var ListResourceRecordSetsCmd = &cli.Command{
	Name: "list-resource-record-sets",
	Flags: []cli.Flag{
		utils.ProfileFlag,
		utils.OutputFlag,
		utils.SortFlag,
		HostedZoneIdFlag,
	},
	Category: "route53",
	Action: func(ctx context.Context, command *cli.Command) error {
		return utils.WithOutput[ResourceRecord](ctx, command, listResourceRecordSets, resourceRecordToTableRow)
	},
}

func listResourceRecordSets(ctx context.Context, command *cli.Command) ([]ResourceRecord, error) {
	client, errClient := utils.GetClient(ctx, command, route53.NewFromConfig)
	if errClient != nil {
		return nil, errClient

	}

	hostedZoneId := command.String("hosted-zone-id")
	if hostedZoneId == "" {
		hostedZones, _ := listHostedZones(ctx, command)
		selectedHz, err := utils.SelectWithFzf(hostedZones, func(hz HostedZone, _ int) string {
			return fmt.Sprintf("%-36s %s", hz.Id, hz.Name)
		})
		if err != nil {
			return nil, err
		}
		fmt.Println("HostedZone:", selectedHz)
		hostedZoneId = strings.Split(selectedHz, " ")[0]
	}

	input := &route53.ListResourceRecordSetsInput{
		HostedZoneId: &hostedZoneId,
	}
	result, err := client.ListResourceRecordSets(ctx, input)
	if err != nil {
		return nil, err
	}

	resourceRecordSets := make([]ResourceRecord, len(result.ResourceRecordSets))
	for i, resource := range result.ResourceRecordSets {

		resourceRecordSets[i] = NewResourceRecord(resource)
	}

	slices.SortFunc(resourceRecordSets, func(a, b ResourceRecord) int {
		return strings.Compare(a.Name, b.Name)
	})
	return resourceRecordSets, nil
}

func resourceRecordToTableRow(hz ResourceRecord) table.Row {
	return table.Row{
		hz.Name,
	}
}
