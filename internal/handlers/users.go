package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kodinggratis/RMNJ.git/domain"
)

type UserHandler struct {
	Service domain.UserService // Dependency Injection
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id")) // Fitur Go 1.22+

	user, _ := h.Service.GetByID(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
