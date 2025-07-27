package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Diegoes7/meetups/models"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON body from JavaScript fetch()
	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	graphqlQuery := `
		mutation Login($email: String!, $password: String!) {
			login(input: {email: $email, password: $password}) {
				authToken {
					accessToken
					expiredAt
				}
				user {
					id
					username
					email
					firstName
					lastName
					createdAt
				}
			}
		}
	`

	variables := map[string]interface{}{
		"email":    loginReq.Email,
		"password": loginReq.Password,
	}

	payload := map[string]interface{}{
		"query":     graphqlQuery,
		"variables": variables,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Could not encode request", http.StatusInternalServerError)
		return
	}

	// resp, err := http.Post("http://localhost:8080/query", "application/json", bytes.NewReader(body))
	graphqlURL := os.Getenv("RENDER_EXTERNAL_URL")
	if graphqlURL == "" {
		graphqlURL = "http://localhost:8080"
	}
	graphqlURL += "/query"

	resp, err := http.Post(graphqlURL, "application/json", bytes.NewReader(body))
	if err != nil {
		http.Error(w, "Failed to send request to GraphQL endpoint", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "GraphQL server error", http.StatusInternalServerError)
		return
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	var graphqlResp models.GraphQLResponse
	if err := json.Unmarshal(respBody, &graphqlResp); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse response: %v", err), http.StatusInternalServerError)
		return
	}

	if graphqlResp.Data.Login.AuthToken == nil || graphqlResp.Data.Login.User == nil {
		http.Error(w, "Invalid login response", http.StatusUnauthorized)
		return
	}

	token := graphqlResp.Data.Login.AuthToken.AccessToken
	user := graphqlResp.Data.Login.User

	http.SetCookie(w, &http.Cookie{
		Name:     "authToken",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		// Secure:   true, // Set to false during local dev if not using HTTPS
		Secure:   os.Getenv("ENV") == "production",
		SameSite: http.SameSiteLaxMode,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		User *models.User `json:"user"`
	}{
		User: user,
	})
}
