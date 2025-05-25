package infrastructure

import (
	"filestorage/models"
	//"storj.io/common/uuid"
	"github.com/google/uuid"
)

type FileRepositoring interface {
	Add(f *models.FileRecord) error
	Get(uuid.UUID) (*models.FileRecord, error)
	//All() (iter.Seq[*models.FileRecord], error)
	All() ([]*models.FileRecord, error)
	//Update(uuid.UUID, *models.FileRecord) error
	//Delete(uuid.UUID) error
}

type RepoError struct {
	msg string
}

func (e *RepoError) Error() string {
	return e.msg
}

func newRepoError() *RepoError {
	return &RepoError{"Bad operation on file repository"}
}

func newRepoErrorWithMsg(msg string) *RepoError {
	return &RepoError{msg}
}
