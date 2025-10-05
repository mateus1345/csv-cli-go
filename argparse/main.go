package argparse

import (
	"csv-cli/csvmetadata"
	"errors"
	"flag"
	"fmt"
	"strings"
)

func Parse() (csvmetadata.CSVMetadata, error) {
	var args csvmetadata.CSVMetadata

	// Define flags with default values
	var delimiter string
	flag.StringVar(&delimiter, "delimiter", ",", "Field delimiter character")
	flag.StringVar(&delimiter, "d", ",", "Field delimiter character (shorthand)")

	var quote string
	flag.StringVar(&quote, "quote", "\"", "Quote character for fields")
	flag.StringVar(&quote, "q", "\"", "Quote character for fields (shorthand)")

	flag.BoolVar(&args.HasHeader, "header", true, "File has header row")
	flag.BoolVar(&args.HasHeader, "H", true, "File has header row (shorthand)")

	flag.StringVar(&args.NullValue, "null", "\\N", "Null value representation")
	flag.StringVar(&args.NullValue, "n", "\\N", "Null value representation (shorthand)")

	flag.Parse()

	var err error
	if args.Delimiter, err = validateDelimiter(delimiter); err != nil {
		return csvmetadata.CSVMetadata{}, err
	}

	if args.Quote, err = validateQuote(quote); err != nil {
		return csvmetadata.CSVMetadata{}, err
	}

	if args.FilePath, err = validateFilePath(); err != nil {
		return csvmetadata.CSVMetadata{}, err
	}
	return args, nil
}

func validateFilePath() (string, error) {
	if flag.NArg() == 0 {
		return "", errors.New("CSV file path is required")
	}
	if flag.NArg() > 1 {
		return "", errors.New("only one CSV file can be processed at a time")
	}

	path := flag.Arg(0)
	if !strings.HasSuffix(strings.ToLower(path), ".csv") {
		return "", errors.New("the file must be in CSV format")
	}
	return path, nil
}

func validateDelimiter(delimiter string) (rune, error) {
	if len(delimiter) != 1 {
		return 0, fmt.Errorf("delimiter must be a single character, got '%s'", delimiter)
	}
	return rune(delimiter[0]), nil
}

func validateQuote(quote string) (rune, error) {
	if len(quote) != 1 {
		return 0, fmt.Errorf("quote must be a single character, got '%s'", quote)
	}
	return rune(quote[0]), nil
}
