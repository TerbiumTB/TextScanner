package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	l *log.Logger
}

func NewHandler(l *log.Logger) *Handler {
	return &Handler{l}
}

//type File struct {
//	filename string
//	content  io.Reader
//}

type UploadResponse struct {
	Filename string `json:"filename"`
}

func (h *Handler) Upload(rw http.ResponseWriter, r *http.Request) {
	vargs := mux.Vars(r)
	filename := vargs["filename"]

	h.l.Printf("Uploading file \"%s\"...", filename)
}

func (h *Handler) Download(rw http.ResponseWriter, r *http.Request) {
	vargs := mux.Vars(r)
	id := vargs["id"]

	h.l.Printf("Downloading file \"%s\"...", id)
}

func (h *Handler) DownloadAll(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Downloading all files...")
}
