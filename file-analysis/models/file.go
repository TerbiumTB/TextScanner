package models

import "github.com/google/uuid"

//type FileOriginalityRecord struct {
//	ID    uuid.UUID `db:"id" json:"id"`
//	Hash  uint32    `db:"hash" json:"hash"`
//	Other uuid.UUID `db:"other" json:"other"`
//	//words   int
//	//symbols int
//}
//
//func NewFileOriginalityRecord(id uuid.UUID, h uint32) *FileOriginalityRecord {
//	return &FileOriginalityRecord{id, h, uuid.Nil}
//}

type FileStat struct {
	Id         uuid.UUID `db:"id" json:"id"`
	Symbols    int       `db:"symbols" json:"symbols"`
	Words      int       `db:"words" json:"words"`
	Sentences  int       `db:"sentences" json:"sentences"`
	Paragraphs int       `db:"paragraphs" json:"paragraphs"`
	Location   string    `db:"location" json:"location"`
}
