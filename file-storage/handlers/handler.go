package handlers

import (
	"filestorage/pkg/json"
	"filestorage/service"
	"github.com/gorilla/mux"
	"io"
	"log"
	"mime/multipart"
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

//type UploadResponse struct {
//	Id string `json:"id"`
//}

func (h *Handler) Upload(wr http.ResponseWriter, r *http.Request) {
	vargs := mux.Vars(r)
	filename := vargs["filename"]
	h.l.Printf("Uploading file \"%s\"...", filename)

	//response := &UploadResponse{}
	var res struct {
		Id string `json:"id"`
	}
	var err error
	//defer r.Body.Close()

	ff, _, _ := r.FormFile("file")
	//io.Copy(os.Stdout, ff)
	res.Id, err = h.s.Upload(filename, ff)

	if err != nil {
		h.l.Printf("[ERROR] Error uploading file \"%s\": %s", filename, err)
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.ToJSON(res, wr)
}

func (h *Handler) Download(wr http.ResponseWriter, r *http.Request) {
	vargs := mux.Vars(r)
	id := vargs["id"]

	h.l.Printf("Downloading file \"%s\"...", id)

	file, err := h.s.Download(id)

	if err != nil {
		h.l.Printf("[ERROR] Error downloading file \"%s\": %s", id, err)
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	mw := multipart.NewWriter(wr)
	defer mw.Close()
	//mw.WriteField("id", id)
	w, _ := mw.CreateFormField("record")
	_ = json.ToJSON(file.Record, w)
	//mw.WriteField("record", file.Record)

	w, _ = mw.CreateFormFile("file", file.Record.Name)

	io.Copy(w, file.Content)
	//var res struct {
	//	Id      string `json:"id"`
	//	Name    string `json:"name"`
	//	Content []byte `json:"content"`
	//}
	//
	//res.Id = id
	////res.Name =
	////_ = json.ToJSON(file.Record, wr)
	////io.Copy(os.Stdout, file.Content)
	////io.Copy(wr, file.Content)
	//json.ToJSON(file, wr)
	//h.l.Println(n)
	//if err != nil {
	//	h.l.Printf("[ERROR] Error downloading file contents \"%s\": %s", id, err)
	//}
	//wr.Write(file.Content)

}

func (h *Handler) GetAllRecords(wr http.ResponseWriter, r *http.Request) {
	h.l.Println("Getting all file records...")

	files, err := h.s.GetAllRecords()

	if err != nil {
		h.l.Printf("[ERROR] Error getting file records: %s", err)
		return
	}

	_ = json.ToJSON(&files, wr)
}

func (h *Handler) GetRecord(wr http.ResponseWriter, r *http.Request) {
	vargs := mux.Vars(r)
	id := vargs["id"]

	h.l.Printf("Getting file \"%s\" record ...", id)

	file, err := h.s.GetRecord(id)
	if err != nil {
		h.l.Printf("[ERROR] Error getting file \"%s\" record: %s", id, err)
		return
	}

	_ = json.ToJSON(file, wr)
}
