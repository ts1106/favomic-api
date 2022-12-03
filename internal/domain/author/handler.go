package author

import (
	"net/http"

	"github.com/ts1106/favomic-api/internal/util/render"
	"github.com/ts1106/favomic-api/internal/util/response"
)

type Handler struct {
	server *Server
}

func NewHandler(s *Server) *Handler {
	return &Handler{server: s}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var param CreateRequest
	param.Decode(r.Body)

	errs := param.Validate()
	if errs != nil {
		code := http.StatusBadRequest
		render.JSON(w, code, response.Errs(code, errs))
		return
	}

	res, err := h.server.Create(r.Context(), param)
	if err != nil {
		code := http.StatusBadRequest
		render.JSON(w, code, response.Err(code, err))
		return
	}
	code := http.StatusCreated
	render.JSON(w, code, res)
}
