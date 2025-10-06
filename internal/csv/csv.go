package csv

import (
	"bufio"
	"csv-cli/internal/csv/parser"
	"csv-cli/internal/models"
	"log"
	"os"
)

func ReadCSV(metadata models.CSVMetadata) models.CSV {
	file, err := os.Open(metadata.FilePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	header := []string{}
	rows := [][]string{}
	columns := make(map[string]models.Column)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values, err := parser.ParseLine(scanner.Text(), metadata)
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
	return models.CSV{
		Header:   header,
		Rows:     rows,
		Columns:  columns,
		Metadata: metadata,
	}
}
