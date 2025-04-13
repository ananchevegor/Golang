package handler

import (
	"example/hello/internal/view/home"
	"net/http"
)

type homeHandler struct{}

func (h homeHandler) handleIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}

func (h homeHandler) handleAbout(w http.ResponseWriter, r *http.Request) error {
	return home.About().Render(r.Context(), w)
}
