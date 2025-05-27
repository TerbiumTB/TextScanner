package handlers

import (
	"filestorage/pkg/json"
	"github.com/gorilla/mux"
	"net/http"
)

type UploadResponse struct {
	ID string `json:"id"`
}

// @Title Upload File
// @Description Загружает файл на сервер и возвращает его уникальный идентификатор
// @Tags Загрузка
// @Accept  multipart/form-data
// @Produce json
// @Param   filename path     string  true  "Имя файла"
// @Param   file     formData file    true  "Файл для загрузки"
// @Success 200      {object} UploadResponse
// @Failure 500      {string} string "Ошибка сервера"
// @Router  /upload/{filename} [post]
func (h *Handler) Upload(wr http.ResponseWriter, r *http.Request) {
	vargs := mux.Vars(r)
	filename := vargs["filename"]
	h.l.Printf("Uploading file \"%s\"...", filename)

	res := UploadResponse{}
	var err error

	ff, _, _ := r.FormFile("file")
	res.ID, err = h.s.Upload(filename, ff)

	if err != nil {
		h.l.Printf("[ERROR] Error uploading file \"%s\": %s", filename, err)
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.ToJSON(res, wr)
}
