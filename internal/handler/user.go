package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/berkeatalay1/rss_go/internal/database"
	"github.com/berkeatalay1/rss_go/internal/utility/models"
	"github.com/berkeatalay1/rss_go/internal/utility/response"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		response.RespondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	response.RespondWithJSON(w, http.StatusOK, models.DatabaseUserToUser(user))
}
