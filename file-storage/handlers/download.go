package handlers

import (
	"filestorage/models"
	"filestorage/pkg/json"
	"github.com/gorilla/mux"
	"io"
	"mime/multipart"
	"net/http"
)

type DownloadResponse struct {
	Record models.FileRecord `json:"record"`
	File   []byte            `json:"-"`
}

// @Title Download File
// @Description Загружает файл по его уникальному идентификатору в формате multipart/form-data
// @Tags Загрузка
// @Produce multipart/form-data
// @Param   id  path  string  true  "Уникальный идентификатор файла"
// @Success 200      {object} DownloadResponse
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /download/{id} [get]
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

	wr.Header().Add("Content-Type", mw.FormDataContentType())
	w, _ := mw.CreateFormField("record")
	_ = json.ToJSON(file.Record, w)

	w, _ = mw.CreateFormFile("file", file.Record.Name)

	io.Copy(w, file.Content)
}
