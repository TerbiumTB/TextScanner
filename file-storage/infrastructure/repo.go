package infrastructure

import (
	"filestorage/models"
	"iter"

	//"github.com/beevik/guid"
	// "golang.org/x/xerrors"
	"storj.io/common/uuid"
)

type FileRepositoring interface {
	Add(f *models.FileRecord) error
	Get(uuid.UUID) (*models.FileRecord, error)
	All() iter.Seq[*models.FileRecord]
	Delete(uuid.UUID) error
}

type FileMap struct {
	files map[uuid.UUID]*models.FileRecord
}

func NewFileMap() *FileMap {
	return &FileMap{make(map[uuid.UUID]*models.FileRecord)}
}

type RepoError struct{}

func (e *RepoError) Error() string {
	return "Bad operation on file repository"
}

func NewRepoError() *RepoError {
	return &RepoError{}
}

type FileNotFoundError struct{}

func (e *FileNotFoundError) Error() string {
	return "file was not found in repository"
}

func NewFileNotFoundError() *FileNotFoundError {
	return &FileNotFoundError{}
}

func (r *FileMap) Get(id uuid.UUID) (*models.FileRecord, error) {
	f, ok := r.files[id]
	if !ok {
		return nil, NewRepoError()
	}
	return f, nil
}

type FileAlreadyExists struct {
	// id uuid.UUID
	// code int
	// msg string
}

func (e *FileAlreadyExists) Error() string {
	return "file with that id already exists in repository"
}

func (r *FileMap) Add(f *models.FileRecord) (err error) {
	f.ID, err = uuid.New()
	if err != nil {
		return err
	}

	if _, found := r.files[f.ID]; found {
		return NewRepoError()
	}

	r.files[f.ID] = f
	return nil
}

func (r *FileMap) All() iter.Seq[*models.FileRecord] {
	return func(yield func(*models.FileRecord) bool) {
		for _, f := range r.files {
			if !yield(f) {
				return
			}
		}
	}
}

func (r *FileMap) Delete(id uuid.UUID) error {
	if _, ok := r.files[id]; !ok {
		return NewRepoError()
	}

	delete(r.files, id)
	return nil
}
