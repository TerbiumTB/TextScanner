package infrastructure

import (
	"fileanalysis/models"
	"github.com/google/uuid"
)

type RepoError struct {
	msg string
}

func (e RepoError) Error() string {
	return e.msg
}

func newRepoError() *RepoError {
	return &RepoError{"Bad operation on repository"}
}

func newRepoErrorWithMsg(msg string) *RepoError {
	return &RepoError{msg}
}

type FileStatRepositoring interface {
	Add(f *models.FileStat) error
	Get(uuid.UUID) (*models.FileStat, error)
	All() ([]*models.FileStat, error)
}
