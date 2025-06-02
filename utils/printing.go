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

	switch command.String("output") {
	case "json":
		PrintJson(list)
	case "table":
		BuildTable(list, itemToTableRow)
	default:
		BuildTable(list, itemToTableRow)
	}

	return nil

}
