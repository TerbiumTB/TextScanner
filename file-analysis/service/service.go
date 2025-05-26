package service

import (
	"bytes"
	//"database/sql"
	"fileanalysis/infrastructure"
	"fileanalysis/models"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type AnalysisNotFoundError struct{}

func (e AnalysisNotFoundError) Error() string { return "file with that ID wasn't analysed before." }

type Serving interface {
	//CheckOriginality(id string) (uuid.UUID, error)
	GetStats(id string) (*models.FileStat, error)
	GetAllStats() ([]*models.FileStat, error)
}

type Service struct {
	client *http.Client
	//r infrastructure.FileOriginalityRepositoring
	stats infrastructure.FileStatRepositoring
}

func NewService(client *http.Client, stats infrastructure.FileStatRepositoring) *Service {
	return &Service{client, stats}
}

func (s *Service) GetStats(id string) (stat *models.FileStat, err error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	if stat, err = s.stats.Get(uid); err == nil {
		log.Println(stat)
		return stat, nil
	}
	log.Println(err)

	stat = &models.FileStat{}
	stat.Id = uid

	content, err := s.downloadFile(uid)
	if err != nil {
		return nil, err
	}

	text := string(content)
	stat.Symbols = len(text)
	stat.Words = len(strings.Fields(text))
	stat.Sentences = strings.Count(text, ". ")
	stat.Paragraphs = strings.Count(text, "\n")

	err = s.stats.Add(stat)
	log.Println(err)

	return stat, nil
}

func (s *Service) GetAllStats() (stats []*models.FileStat, err error) {
	return s.stats.All()
}

//func (s *Service) CheckOriginality(id string) (uuid.UUID, error) {
//	uid, err := uuid.Parse(id)
//	if err != nil {
//		return uuid.Nil, err
//	}
//
//	var content []byte
//
//	r, err := s.r.Get(uid)
//	log.Printf("%#v", r)
//
//	//content := make([]byte, 50)
//
//	log.Println("here")
//	switch err.(type) {
//	case nil:
//		log.Println("there")
//		content, err = s.downloadFile(uid)
//		log.Printf("%#v\n", content)
//		if err != nil {
//			return uuid.Nil, err
//		}
//
//	case infrastructure.RepoError:
//		log.Println("over here")
//		content, err = s.downloadFile(uid)
//		log.Printf("%#v\n", content)
//
//		if err != nil {
//			return uuid.Nil, err
//		}
//		h, err := calculateHash(content)
//		if err != nil {
//			return uuid.Nil, err
//		}
//		r = models.NewFileOriginalityRecord(uid, h)
//
//		//s.r.Add()
//
//	default:
//		log.Println("over here")
//		return uuid.Nil, err
//	}
//
//	fmt.Printf("%#v\n", r)
//	//content
//	defer s.r.Add(r)
//
//	res, err := s.getRequest("/record")
//	var others []struct {
//		ID       uuid.UUID `json:"id"`
//		Filename string    `json:"filename"`
//		Location string    `json:"location"`
//	}
//
//	err = json.FromJSON(others, res.Body)
//	if err != nil {
//		return uuid.Nil, err
//	}
//
//	for _, o := range others {
//
//		oo, err := s.r.Get(o.ID)
//
//		if err != nil {
//			ocontent, err := s.downloadFile(o.ID)
//			if err != nil {
//				continue
//			}
//			oh, _ := calculateHash(ocontent)
//
//			oo = models.NewFileOriginalityRecord(o.ID, oh)
//		}
//		if oo.Hash == r.Hash {
//			ocontent, err := s.downloadFile(oo.ID)
//			if err != nil {
//				continue
//			}
//			if slices.Equal(content, ocontent) {
//				r.Other = oo.ID
//				return oo.ID, nil
//			}
//
//		}
//	}
//
//	return uuid.Nil, nil
//

//func calculateHash(b []byte) (uint32, error) {
//	h := crc32.NewIEEE()
//	_, err := h.Write(b)
//
//	if err != nil {
//		return 0, err
//	}
//	return h.Sum32(), nil
//}

func (s *Service) downloadFile(id uuid.UUID) ([]byte, error) {
	log.Printf("Downloading file with ID: %v", id)
	resp, err := s.getRequest("/download/", id)
	if err != nil {
		return nil, err
	}

	contentType := resp.Header.Get("Content-Type")
	_, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to parse content type: %v", err)
	}
	boundary := params["boundary"]
	if boundary == "" {
		return nil, fmt.Errorf("no boundary in content type")
	}

	mr := multipart.NewReader(resp.Body, boundary)

	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if part.FormName() == "file" {
			var buf bytes.Buffer
			if _, err := io.Copy(&buf, part); err != nil {
				return nil, fmt.Errorf("error reading file content: %v", err)
			}
			return buf.Bytes(), nil
		}
	}

	return nil, AnalysisNotFoundError{}
}

func (s *Service) getRequest(request string, id uuid.UUID) (*http.Response, error) {
	root := os.Getenv("FILE_STORAGE_URL")
	//log.Println(root + fmt.Sprintf(request, a))
	return s.client.Get(root + request + id.String())
}
