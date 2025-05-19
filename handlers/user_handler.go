package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Diegoes7/meetups/domain"
)

// UsersHandler returns all registered users as JSON
func UsersHandler(d *domain.Domain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := d.UserRepo.GetUsers()
		if err != nil {
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}

		// Optional: filter fields if needed (e.g., exclude passwords, etc.)
		type userResponse struct {
			ID   string  `json:"id"`
			Name string `json:"name"`
		}
		response := make([]userResponse, len(users))
		for i, u := range users {
			response[i] = userResponse{ID: u.ID, Name: u.FirstName + " " + u.LastName}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
