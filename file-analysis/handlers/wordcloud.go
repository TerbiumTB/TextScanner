package handlers

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

// @Title File Word Cloud
// @Summary Получить облако слов
// @Description Генерирует и возвращает облако слов для файла в формате PNG
// @Tags Облако слов
// @Produce png
// @Param id path string true "ID файла"
// @Success 200 {file} binary "Изображение PNG с облаком слов"
// @Header 200 {string} Content-Type "image/png"
// @Failure 500 {string} string "Ошибка генерации облака слов"
// @Router /wordcloud/{id} [get]
func (h *Handler) GetWordCloud(w http.ResponseWriter, r *http.Request) {
	vargs := mux.Vars(r)
	id := vargs["id"]

	img, err := h.s.GetWordCloud(id)
	if err != nil {
		h.l.Printf("[ERROR] Error getting word cloud for file \"%s\": %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer img.Close()

	w.Header().Set("Content-Type", "image/png")
	_, _ = io.Copy(w, img)
}
