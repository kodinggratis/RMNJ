package routers

import (
	"net/http"

	"github.com/kodinggratis/RMNJ.git/internal/handlers"
	"github.com/kodinggratis/RMNJ.git/internal/middleware"
)

type RouterDeps struct {
	UserH *handlers.UserHandler
}

func NewRouter(deps RouterDeps) http.Handler {
	mux := http.NewServeMux()
	RegisterUserRoutes(mux, deps.UserH)

	// Bungkus dengan Middleware (Logger & Recovery)
	return middleware.Recovery(mux)
}
