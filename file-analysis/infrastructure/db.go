package infrastructure

import (
	"database/sql"
	"errors"
	"fileanalysis/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type FileStatDBX struct {
	db *sqlx.DB
}

func NewFileStatsDBX(db *sqlx.DB) (*FileStatDBX, error) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS stats (
			id UUID PRIMARY KEY,
			symbols BIGINT,
			words BIGINT,
			sentences BIGINT,
			paragraphs BIGINT
		)
	`)

	if err != nil {
		return nil, err
	}
	return &FileStatDBX{db}, nil
}

func (f *FileStatDBX) Add(stat *models.FileStat) (err error) {
	_, err = f.db.NamedExec(`
			INSERT INTO stats (id, symbols, words, sentences, paragraphs ) 
			VALUES (:id, :symbols, :words, :sentences, :paragraphs)
			`, stat)

	return
}

func (f *FileStatDBX) Get(id uuid.UUID) (stat *models.FileStat, err error) {
	stat = &models.FileStat{}
	err = f.db.Get(stat, `SELECT * FROM stats WHERE id = $1`, id)
	//log.Println(stat)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, newRepoError()
	}
	return
}

func (f *FileStatDBX) All() (stats []*models.FileStat, err error) {
	err = f.db.Select(&stats, `SELECT * FROM stats`)

	return

}
