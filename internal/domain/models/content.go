package models

type ContentType string

const (
	Image ContentType = "image"
	Video ContentType = "video"
)

type ContentMetadata struct {
	Type     ContentType
	MimeType string
	Size     int64
}
