package handlers

import (
	"filestorage/pkg/json"
	"filestorage/service"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Handler struct {
	l *log.Logger
	s service.Serving
}

func NewHandler(l *log.Logger, s service.Serving) *Handler {
	return &Handler{l, s}
}

//type File struct {
//	filename string
//	content  io.Reader
//}

type UploadResponse struct {
	Id string `json:"id"`
}

func (h *Handler) Upload(wr http.ResponseWriter, r *http.Request) {
	vargs := mux.Vars(r)
	filename := vargs["filename"]
	h.l.Printf("Uploading file \"%s\"...", filename)

	response := &UploadResponse{}
	var err error
	//defer r.Body.Close()

	ff, _, _ := r.FormFile("file")
	//io.Copy(os.Stdout, ff)
	response.Id, err = h.s.Add(filename, ff)

	if err != nil {
		h.l.Printf("[ERROR] Error uploading file \"%s\": %s", filename, err)
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.ToJSON(response, wr)
}

func (h *Handler) Download(wr http.ResponseWriter, r *http.Request) {
	vargs := mux.Vars(r)
	id := vargs["id"]

	h.l.Printf("Downloading file \"%s\"...", id)

	file, err := h.s.Get(id)

	h.l.Printf("Downloading file %#v...", file)

	if err != nil {
		h.l.Printf("[ERROR] Error downloading file \"%s\": %s", id, err)
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.ToJSON(file.Record, wr)
	//io.Copy(os.Stdout, file.Content)
	io.Copy(wr, file.Content)

	//h.l.Println(n)
	//if err != nil {
	//	h.l.Printf("[ERROR] Error downloading file contents \"%s\": %s", id, err)
	//}
	//wr.Write(file.Content)

}

func (h *Handler) DownloadAll(wr http.ResponseWriter, r *http.Request) {
	h.l.Println("Downloading all file records...")

	files, err := h.s.All()

	if err != nil {
		h.l.Printf("[ERROR] Error downloading all file records: %s", err)
		return
	}

	_ = json.ToJSON(&files, wr)
}
