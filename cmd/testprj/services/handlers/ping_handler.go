package handlers

import (
	"net/http"
	"github.com/ducnt114/testprj/utils"
)

type PingHandler struct {
}

type pingResponse struct {
	Success bool `json:"success"`
}

func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJSON(w, &pingResponse{Success: true})
}
