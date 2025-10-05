package main

import (
	"csv-cli/argparse"
	"csv-cli/csv"
	"fmt"
)

func main() {
	metadata, err := argparse.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}
	f := csv.ReadCSV(metadata)
	f.PrintStatistics()
}
