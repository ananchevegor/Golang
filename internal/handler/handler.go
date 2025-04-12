package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handlerFunc func(http.ResponseWriter, *http.Request) error

type Dependencies struct {
	AssetsFS http.FileSystem
}

func RegisterRoutes(r *chi.Mux, deps Dependencies) {

	home := homeHandler{}

	r.Get("/", handler(home.handleIndex))

	r.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(deps.AssetsFS)))
}

func handler(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			handlerError(w, r, err)
		}
	}
}

func handlerError(w http.ResponseWriter, r *http.Request, err error) {
	slog.Error("error during request", slog.String("err", err.Error()))
}
