package models

import "io"

type ContentType string

const (
	Image ContentType = "image"
	Video ContentType = "video"
)

type ContentMetadata struct {
	URL      string
	File     io.Reader
	Type     ContentType
	MimeType string
	Size     int64
}
