package csvparser

import (
	"csv-cli/csvmetadata"
	"strings"
)

func ParseLine(line string, metadata csvmetadata.CSVMetadata) ([]string, error) {
	values := make([]string, 0)
	ignoreDelimiter := false
	valueStart := 0
	for i, r := range line {
		if r == metadata.Delimiter && !ignoreDelimiter {
			value := line[valueStart:i]
			values = append(values, trimQuotes(value))
			valueStart = i + 1
		} else if r == metadata.Quote {
			ignoreDelimiter = !ignoreDelimiter
		}
	}
	lastValue := line[valueStart:]
	values = append(values, trimQuotes(lastValue))
	return values, nil
}

func trimQuotes(value string) string {
	return strings.Trim(value, `"`)
}
