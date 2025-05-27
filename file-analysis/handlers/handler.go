package handlers

import (
	"fileanalysis/service"
	"log"
	"net/http"
)

//var storageClient *http.Client

type Handler struct {
	l *log.Logger
	c *http.Client
	s service.Serving
}

func NewHandler(l *log.Logger, c *http.Client, s service.Serving) *Handler {
	return &Handler{l, c, s}
}

//type AnalysisResponse struct {
//	Filename string `json:"filename"`
//}

//func (h *Handler) CheckOriginalityHandler(wr http.ResponseWriter, r *http.Request) {
//	//var res struct{
//	//	Filename string `json:"filename"`
//	//}
//	vargs := mux.Vars(r)
//	id := vargs["id"]
//
//	h.l.Printf("Analysing file \"%s\"...", id)
//	//h.l.Printf("djfhashfjkadhfjkashdfjkhdafgjkhdafjkghdjksfghkjsdfhgjkdfhgkjsdfgjksdfhgj")
//
//	uid, err := h.s.CheckOriginality(id)
//
//	if err != nil {
//		h.l.Printf("[ERROR] Error analysing file \"%s\" from storage: %v", id, err)
//		http.Error(wr, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	var resp struct {
//		ID uuid.UUID `json:"id"`
//	}
//	resp.ID = uid
//
//	//
//	//if err != nil {
//	//	h.l.Printf("[ERROR] could not download file %s", id)
//	//	return
//	//}
//	//
//	_ = json.ToJSON(resp, wr)
//}
