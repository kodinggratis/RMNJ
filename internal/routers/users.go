package routers

import (
	"net/http"

	"github.com/kodinggratis/RMNJ.git/internal/handlers"
)

func RegisterUserRoutes(mux *http.ServeMux, h *handlers.UserHandler) {
	mux.HandleFunc("GET /v1/users/{id}", h.GetUser)
}
