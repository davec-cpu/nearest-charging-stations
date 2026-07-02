package handler

import (
	"nearest-charging-stations/templates"
	"net/http"
)

type HomeHandler struct {
}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Home(
	w http.ResponseWriter,
	r *http.Request,
) {
	err := templates.Home().Render(r.Context(), w)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
	}
}
