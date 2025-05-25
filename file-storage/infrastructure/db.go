package infrastructure

import (
	"filestorage/models"
	"github.com/jmoiron/sqlx"
	//"storj.io/common/uuid"
	"github.com/google/uuid"
)

type FileDBX struct {
	db *sqlx.DB
}

func NewFileDBX(db *sqlx.DB) (*FileDBX, error) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS files (
			id UUID PRIMARY KEY,
			name TEXT NOT NULL,
			location TEXT NOT NULL
		)
	`)

	if err != nil {
		return nil, err
	}

	return &FileDBX{db}, nil
}

func (fdb *FileDBX) Get(id uuid.UUID) (f *models.FileRecord, err error) {
	f = &models.FileRecord{}
	err = fdb.db.Get(f, `SELECT * FROM files WHERE id = $1`, id)

	if err != nil {
		return nil, err
	}

	return
}

func (fdb *FileDBX) Add(f *models.FileRecord) (err error) {
	_, err = fdb.db.NamedExec(`INSERT INTO files (id, name, location) VALUES (:id, :name, :location)`, f)

	return
}

//func (fdb *FileDBX) All() (iter.Seq[*models.FileRecord], error) {
//	f := &[]models.FileRecord{}
//	err = fdb.db.Select(f, `SELECT * FROM files`)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return f, err
//}

func (fdb *FileDBX) All() (f []*models.FileRecord, err error) {
	err = fdb.db.Select(&f, `SELECT * FROM files`)

	if err != nil {
		return nil, err
	}

	return
}
