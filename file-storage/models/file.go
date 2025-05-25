package models

import (
	"io"
	"storj.io/common/uuid"
)

type FileRecord struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
}

func NewFileRecord(name string, location string) *FileRecord {
	return &FileRecord{
		Name:     name,
		Location: location,
	}
}

type File struct {
	Record  *FileRecord `json:"record"`
	Content io.Reader   `json:"content"`
}
