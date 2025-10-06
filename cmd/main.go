package main

import (
	"csv-cli/internal/cli"
	"csv-cli/internal/csv"
	"fmt"
)

func main() {
	metadata, err := cli.ParseArgs()
	if err != nil {
		fmt.Println(err)
		return
	}
	f := csv.ReadCSV(metadata)
	f.PrintStatistics()
}
