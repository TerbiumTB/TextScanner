package handlers

import (
	"filestorage/service"
	"log"
)

type Handler struct {
	l *log.Logger
	s service.Serving
}

func NewHandler(l *log.Logger, s service.Serving) *Handler {
	return &Handler{l, s}
}
