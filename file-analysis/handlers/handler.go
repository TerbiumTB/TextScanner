package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//var storageClient *http.Client

func init() {

}

type Handler struct {
	l *log.Logger
	c *http.Client
}

func NewHandler(l *log.Logger, c *http.Client) *Handler {
	return &Handler{l, c}
}

type AnalysisResponse struct {
	Filename string `json:"filename"`
}

func (h *Handler) Analyse(wr http.ResponseWriter, r *http.Request) {
	//var res struct{
	//	Filename string `json:"filename"`
	//}
	vargs := mux.Vars(r)
	id := vargs["id"]

	h.l.Printf("Analysing file \"%s\"...", id)

	//sr, err := h.c.Download(fmt.Sprintf("/download/%s", id))
	//
	//if err != nil {
	//	h.l.Printf("[ERROR] could not download file %s", id)
	//	return
	//}
	//
	//_ = pkg.ToJSON(sr.Body, wr)
}
