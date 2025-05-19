package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Diegoes7/meetups/domain"
)

type InviteRequest struct {
	MeetupID string `json:"meetupId"`
	UserID   string `json:"userId"`
}

func InviteUserHandler(domain *domain.Domain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req InviteRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}

		_, err := domain.InvitationRepo.InviteUser(req.MeetupID, req.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invited"))
	}
}
