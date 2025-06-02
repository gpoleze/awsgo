package utils

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"reflect"
	"strings"
)

type getRowFromItem[T any] func(T) table.Row

type BuildTableParams[T any] struct {
	ListOfItems       []T
	ItemToRowFunction getRowFromItem[T]
	Header            []string
}

func BuildTableWithHeader[T any](params BuildTableParams[T], sortBy []table.SortBy) {
	var header table.Row

	if params.Header == nil {
		for _, field := range reflect.VisibleFields(reflect.TypeOf(params.ListOfItems[0])) {
			header = append(header, strings.ToUpper(field.Name))
		}
	} else {
		for _, field := range params.Header {
			header = append(header, field)
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(header)

	for _, item := range params.ListOfItems {
		t.AppendRow(params.ItemToRowFunction(item))
	}

	t.SortBy(sortBy)
	t.Render()
	return
}

func BuildTable[T any](listOfItems []T, ItemToRowFunction getRowFromItem[T]) {
	params := BuildTableParams[T]{listOfItems, ItemToRowFunction, nil}
	BuildTableWithHeader(params, nil)
	return
}

func BuildTableSortedBy[T any](listOfItems []T, ItemToRowFunction getRowFromItem[T], sortBy []table.SortBy) {
	params := BuildTableParams[T]{listOfItems, ItemToRowFunction, nil}
	BuildTableWithHeader(params, sortBy)
	return
}
