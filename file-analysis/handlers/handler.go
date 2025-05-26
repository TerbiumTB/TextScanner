package handlers

import (
	"fileanalysis/pkg/json"
	"fileanalysis/service"
	"github.com/gorilla/mux"
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

type AnalysisResponse struct {
	Filename string `json:"filename"`
}

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

func (h *Handler) GetStatsHandler(wr http.ResponseWriter, r *http.Request) {
	vargs := mux.Vars(r)
	id := vargs["id"]

	h.l.Printf("Analysing file \"%s\"...", id)

	stat, err := h.s.GetStats(id)
	if err != nil {
		h.l.Printf("[ERROR] Error getting stats for file \"%s\": %v", id, err)
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.ToJSON(stat, wr)
}

func (h *Handler) GetAllStatsHandler(wr http.ResponseWriter, r *http.Request) {
	stats, err := h.s.GetAllStats()

	if err != nil {
		h.l.Printf("[ERROR] Error getting all stats: %v", err)
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.ToJSON(stats, wr)
}
