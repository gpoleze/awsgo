package ses

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v3"
	"gitlab.com/gabriel.poleze/awsgo/utils"
)

var startDateFlag = &cli.StringFlag{
	Name:  "start-date",
	Usage: "Used to filter the list of suppressed email destinations so that it only includes addresses that were added to the list after a specific date.",
}

var DescribeInstancesCmd = &cli.Command{
	Name:                  "list-suppressed-destinations",
	Usage:                 "Retrieves  a  list  of email addresses that are on the suppression list\n       for your account..",
	EnableShellCompletion: true,
	Flags: []cli.Flag{
		utils.RegionFlag,
		utils.ProfileFlag,
		utils.OutputFlag,
		utils.SortFlag,
		utils.SortOrderFlag,
		startDateFlag,
	},
	Category: "ses",
	Action: func(ctx context.Context, command *cli.Command) error {
		_ = utils.SortFlag.Set("Value", "email_address")
		_ = utils.SortOrderFlag.Set("Value", "asc")
		return utils.WithOutput[types.SuppressedDestinationSummary](ctx, command, listSuppressedDestinations, listItemToTableRow)
	},
}

func listSuppressedDestinations(ctx context.Context, command *cli.Command) ([]types.SuppressedDestinationSummary, error) {
	client, errClient := utils.GetClient(ctx, command, sesv2.NewFromConfig)
	if errClient != nil {
		return nil, errClient
	}

	var pageSize int32 = 1000
	startDate := time.Now().AddDate(0, 0, -1)
	startDateString := command.String("start-date")
	if startDateString != "" {
		startDate, _ = time.Parse("2006-01-02", startDateString)
	}
	input := &sesv2.ListSuppressedDestinationsInput{
		PageSize:  &pageSize,
		StartDate: &startDate,
	}

	result, err := client.ListSuppressedDestinations(ctx, input)
	if err != nil {
		return nil, err
	}

	emails := result.SuppressedDestinationSummaries

	for result.NextToken != nil {
		input.NextToken = result.NextToken
		result, err = client.ListSuppressedDestinations(ctx, input)
		if err != nil {
			return nil, err
		}
		emails = append(emails, result.SuppressedDestinationSummaries...)
	}

	return emails, nil
}

func listItemToTableRow(s types.SuppressedDestinationSummary) table.Row {
	return table.Row{
		*s.EmailAddress,
		*s.LastUpdateTime,
		s.Reason,
	}
}
