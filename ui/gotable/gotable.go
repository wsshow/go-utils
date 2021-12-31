package gotable

import (
	"fmt"
	"github.com/liushuochen/gotable"
)

func WorkFlow() {
	table, err := gotable.Create("version", "description")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	_ = table.AddRow([]string{"gotable 5", "Safe: New table type to enhance concurrency security"})
	_ = table.AddRow([]string{"gotable 4", "Colored: Print colored column"})
	_ = table.AddRow([]string{"gotable 3", "Storage: Store the table data as a file"})
	_ = table.AddRow([]string{"gotable 2", "Simple: Use simpler APIs to control table"})
	_ = table.AddRow([]string{"gotable 1", "Gotable: Print a beautiful ASCII table"})

	fmt.Println(table)
}
