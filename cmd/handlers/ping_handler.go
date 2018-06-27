package handlers

import (
	"github.com/ducnt114/testprj/utils"
	"net/http"
)

type PingHandler struct {
}

type pingResponse struct {
	Success bool `json:"success"`
}

func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJSON(w, &pingResponse{Success: true})
}
