package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v3"
	"strings"
)

func PrintJson[T any](item T) {
	val, _ := json.MarshalIndent(item, "", "    ")
	fmt.Println(string(val))
}

func WithOutput[T any](
	ctx context.Context,
	command *cli.Command,
	function func(ctx context.Context, command *cli.Command) ([]T, error),
	itemToTableRow func(T) table.Row,
) error {
	list, err := function(ctx, command)

	if err != nil {
		return err
	}

	var sortBy []table.SortBy
	if columnsToSort := command.StringSlice(SortFlag.Name); columnsToSort != nil && len(columnsToSort) != 0 {
		for _, col := range columnsToSort {
			sortBy = append(sortBy, table.SortBy{
				Name:       strings.ToUpper(col),
				IgnoreCase: true,
			})
		}
	}

	switch command.String("output") {
	case "json":
		PrintJson(list)
	case "table":
		BuildTableSortedBy(list, itemToTableRow, sortBy)
	default:
		BuildTableSortedBy(list, itemToTableRow, sortBy)
	}

	return nil

}
