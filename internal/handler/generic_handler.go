package handler

import (
	"net/http"

	"github.com/berkeatalay1/rss_go/internal/utility/response"
)

func Readiness(w http.ResponseWriter, r *http.Request) {
	response.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func Err(w http.ResponseWriter, r *http.Request) {
	response.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
