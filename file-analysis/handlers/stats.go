package handlers

import (
	"fileanalysis/pkg/json"
	"github.com/gorilla/mux"
	"net/http"
)

// @Title All File Stats
// @Summary Получить всю статистику
// @Description Возвращает статистику по всем файлам
// @Tags Статистика
// @Produce json
// @Success 200 {array} models.FileStat "Список статистики"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /stats [get]
func (h *Handler) GetAllStatsHandler(wr http.ResponseWriter, r *http.Request) {
	stats, err := h.s.GetAllStats()

	if err != nil {
		h.l.Printf("[ERROR] Error getting all stats: %v", err)
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.ToJSON(stats, wr)
}

// @Title File stats
// @Summary Получить статистику файла
// @Description Возвращает статистику по файлу
// @Tags Статистика
// @Produce json
// @Param id path string true "ID файла"
// @Success 200 {object} models.FileStat "Статистика файла"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /stats/{id} [get]
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
