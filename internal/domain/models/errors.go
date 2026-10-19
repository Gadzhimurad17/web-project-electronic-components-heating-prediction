package models

import "errors"

var (
	ErrS3Operation      = errors.New("storage operation failed")
	ErrFileTooLarge     = errors.New("file size exceeds allowed limit")
	ErrFileEmpty        = errors.New("File is empty")
	ErrFileNotAccesible = errors.New("Can`t reach file")
	ErrInvalidFileType  = errors.New("Invalid File type")
)
