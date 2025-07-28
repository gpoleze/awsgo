package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/urfave/cli/v3"
)

func PrintJson[T any](item T) {
	val, _ := json.MarshalIndent(item, "", "    ")
	fmt.Println(string(val))
}

func WithTableOutput[T any](
	ctx context.Context,
	command *cli.Command,
	function func(ctx context.Context, command *cli.Command) ([]T, error),
	tableParams BuildTableParams[T],
) error {
	list, err := function(ctx, command)
	if err != nil {
		return err
	}

	tableParams.ListOfItems = list

	var sortBy []table.SortBy
	if columnsToSort := command.StringSlice(SortFlag.Name); columnsToSort != nil && len(columnsToSort) != 0 {

		mode := table.Asc
		if command.String(SortOrderFlag.Name) == "dsc" {
			mode = table.Dsc
		}

		for _, col := range columnsToSort {
			tableParams.SortBy = append(sortBy, table.SortBy{
				Name:       col,
				IgnoreCase: true,
				Mode:       mode,
			})
		}
	}

	switch command.String("output") {
	case "json":
		PrintJson(list)
	case "table":
		BuildTable[T](tableParams)
	default:
		BuildTable[T](tableParams)
	}

	return nil

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

		mode := table.Asc
		if command.String(SortOrderFlag.Name) == "dsc" {
			mode = table.Dsc
		}

		for _, col := range columnsToSort {
			sortBy = append(sortBy, table.SortBy{
				Name:       col,
				IgnoreCase: true,
				Mode:       mode,
			})
		}
	}

	switch command.String("output") {
	case "json":
		PrintJson(list)
	case "table":
		BuildTable[T](BuildTableParams[T]{list,
			itemToTableRow,
			nil,
			sortBy,
			"",
		})
	default:
		BuildTableSortedBy(list, itemToTableRow, sortBy)
	}

	return nil

}
