package infrastructure

import (
	"filestorage/models"
	"maps"
	"slices"
	//"storj.io/common/uuid"
	"github.com/google/uuid"
)

type FileMap struct {
	files map[uuid.UUID]*models.FileRecord
}

func NewFileMap() *FileMap {
	return &FileMap{make(map[uuid.UUID]*models.FileRecord)}
}

func (m *FileMap) Get(id uuid.UUID) (*models.FileRecord, error) {
	f, ok := m.files[id]
	if !ok {
		return nil, newRepoError()
	}
	return f, nil
}

func (m *FileMap) Add(f *models.FileRecord) (err error) {
	//f.ID, err = uuid.New()
	//if err != nil {
	//	return err
	//}

	if _, found := m.files[f.ID]; found {
		return newRepoError()
	}

	m.files[f.ID] = f
	return nil
}

//func (r *FileMap) All() iter.Seq[*models.FileRecord] {
//	return func(yield func(*models.FileRecord) bool) {
//		for _, f := range r.files {
//			if !yield(f) {
//				return
//			}
//		}
//	}
//}

//func (m *FileMap) All() (iter.Seq[*models.FileRecord], error) {
//	return maps.Values(m.files), nil
//}

func (m *FileMap) All() ([]*models.FileRecord, error) {
	f := slices.Collect(maps.Values(m.files))

	return f, nil
}

func (m *FileMap) Delete(id uuid.UUID) error {
	if _, ok := m.files[id]; !ok {
		return newRepoError()
	}

	delete(m.files, id)
	return nil
}

func (m *FileMap) Update(id uuid.UUID, f *models.FileRecord) error {
	if _, ok := m.files[id]; !ok {
		return newRepoError()
	}

	m.files[id] = f
	return nil
}
