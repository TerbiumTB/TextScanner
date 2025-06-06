package service

import (
	"filestorage/models"
	//"storj.io/common/uuid"
	"github.com/google/uuid"
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
	Upload(filename string, content io.ReadCloser) (string, error)
	Download(id string) (*models.File, error)
	GetRecord(id string) (*models.FileRecord, error)
	GetAllRecords() ([]*models.FileRecord, error)
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

func (s *Service) Upload(filename string, content io.ReadCloser) (string, error) {
	id := uuid.New()

	//path.Join(id.String(), filename)
	location, err := s.storage.Save(id.String(), content)

	if err != nil {

		return "", err
	}

	file := models.NewFileRecord(id, filename, location)

	err = s.repo.Add(file)
	if err != nil {
		return "", err
	}

	return file.ID.String(), nil
}

func (s *Service) Download(id string) (f *models.File, err error) {
	uid, err := uuid.Parse(id)

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

func (s *Service) GetRecord(id string) (*models.FileRecord, error) {
	uid, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	return s.repo.Get(uid)
}

func (s *Service) GetAllRecords() (f []*models.FileRecord, err error) {
	return s.repo.All()
	//var a []*models.File
	//it, err := s.repo.GetAllRecords()
	//if err != nil {
	//	return nil, err
	//}
	//return slices.Collect(it), nil
}
