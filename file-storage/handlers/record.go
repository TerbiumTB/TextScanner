package handlers

import (
	"filestorage/pkg/json"
	"github.com/gorilla/mux"
	"net/http"
)

// @Title Get All Records Records
// @Description Получает метаданные всех файлов
// @Tags Инфо
// @Produce json
// @Success 200 {array} models.FileRecord "Список метаданных файлов"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /record [get]
func (h *Handler) GetAllRecords(wr http.ResponseWriter, r *http.Request) {
	h.l.Println("Getting all file records...")

	files, err := h.s.GetAllRecords()

	if err != nil {
		h.l.Printf("[ERROR] Error getting file records: %s", err)
		return
	}

	_ = json.ToJSON(&files, wr)
}

// @Title Get File Record
// @Description Получает метаданные файла по ID
// @Tags Инфо
// @Produce json
// @Param   id  path  string  true  "Уникальный идентификатор файла"
// @Success 200 {object} models.FileRecord "Метаданные файла"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /record/{id} [get]
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
