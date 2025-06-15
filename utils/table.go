package utils

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"reflect"
)

type getRowFromItem[T any] func(T) table.Row

type BuildTableParams[T any] struct {
	ListOfItems       []T
	ItemToRowFunction getRowFromItem[T]
	Header            []string
	SortBy            []table.SortBy
	Title             string
}

func BuildTable[T any](params BuildTableParams[T]) {
	if len(params.ListOfItems) == 0 {
		fmt.Println("[]")
		return
	}

	var header table.Row

	if params.Header == nil {
		for _, field := range reflect.VisibleFields(reflect.TypeOf(params.ListOfItems[0])) {
			fieldName := strcase.ToSnake(field.Name)
			header = append(header, fieldName)
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

	t.SortBy(params.SortBy)
	t.SetTitle(params.Title)
	t.Render()
	return
}

func BuildTableSortedBy[T any](listOfItems []T, ItemToRowFunction getRowFromItem[T], sortBy []table.SortBy) {
	params := BuildTableParams[T]{
		listOfItems,
		ItemToRowFunction,
		nil,
		sortBy,
		reflect.TypeOf(listOfItems).Elem().Name(),
	}
	BuildTable(params)
	return
}
