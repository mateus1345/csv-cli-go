package csv

import (
	"bufio"
	"csv-cli/csvmetadata"
	"csv-cli/csvparser"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Column []string

type CSV struct {
	Header   []string
	Rows     [][]string
	Columns  map[string]Column
	Metadata csvmetadata.CSVMetadata
}

func (csv CSV) Mean(column string) float64 {
	var sum float64
	var count int
	for _, value := range csv.Columns[column] {
		if value == csv.Metadata.NullValue {
			continue
		}
		valueFloat, err := strconv.ParseFloat(value, 64)
		if err != nil {
			continue
		}
		sum += valueFloat
		count++
	}
	if count == 0 {
		return 0
	}
	return sum / float64(count)
}

func (csv CSV) Mode(column string) string {
	frequency := make(map[string]int)
	for _, value := range csv.Columns[column] {
		if value == csv.Metadata.NullValue {
			continue
		}
		frequency[value]++
	}
	var mode string
	var maxCount int
	for value, count := range frequency {
		if count > maxCount {
			maxCount = count
			mode = value
		}
	}
	return mode
}

func (csv CSV) PrintStatistics() {
	for _, column := range csv.Header {
		fmt.Printf("Column: %s, Mode: %s, Mean: %.2f\n", column, csv.Mode(column), csv.Mean(column))
	}
}

func ReadCSV(metadata csvmetadata.CSVMetadata) CSV {
	file, err := os.Open(metadata.FilePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	header := []string{}
	rows := [][]string{}
	columns := make(map[string]Column)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values, err := csvparser.ParseLine(scanner.Text(), metadata)
		if err != nil {
			log.Printf("Error parsing line: %v", err)
			continue
		}

		if len(header) == 0 && metadata.HasHeader {
			header = values
		} else {
			rows = append(rows, values)
			for i, column := range header {
				columns[column] = append(columns[column], values[i])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error during scanning: %v", err)
	}
	return CSV{header, rows, columns, metadata}
}
