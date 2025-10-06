package models

type CSVMetadata struct {
	FilePath  string
	Delimiter rune
	Quote     rune
	NullValue string
	HasHeader bool
}
