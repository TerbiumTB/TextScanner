package service

import (
	"filestorage/models"
	"slices"
	"storj.io/common/uuid"

	//"filestorage/infrastructure"
	"filestorage/infrastructure"

	// "github.com/beevik/guid"
	// "hash/crc32"
	"io"
	"sync"
	// "golang.org/x/mod/sumdb/storage"
	// "golang.org/x/xerrors"
)

type Serving interface {
	Add(filename string, content io.Reader) (string, error)
	Get(id string) (*models.File, error)
	All() []*models.FileRecord
}

type Service struct {
	repo    infrastructure.FileRepositoring
	storage infrastructure.FileStorer
}

var once sync.Once

func NewService(r infrastructure.FileRepositoring, s infrastructure.FileStorer) *Service {
	// once.Do(

	// )
	return &Service{
		repo:    r,
		storage: s,
	}
}

func (s *Service) Add(filename string, content io.Reader) (string, error) {
	location, err := s.storage.Save(filename, content)
	if err != nil {
		return "", err
	}

	file := models.NewFileRecord(filename, location)

	err = s.repo.Add(file)
	if err != nil {
		return "", err
	}

	return file.ID.String(), nil
}

func (s *Service) Get(id string) (f *models.File, err error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}

	f = &models.File{}

	f.Record, err = s.repo.Get(uid)

	if err != nil {
		return nil, err
	}

	f.Content, err = s.storage.Load(f.Record.Location)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) All() []*models.FileRecord {
	//var a []*models.File
	return slices.Collect(s.repo.All())
}
