package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Diegoes7/meetups/models"
)

func MeHandler(w http.ResponseWriter, r *http.Request) {
	// Check for the authToken in cookies
	cookie, err := r.Cookie("authToken")
	if err != nil {
		// No cookie — return empty object
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
		return
	}

	token := cookie.Value

	// GraphQL query to fetch user info
	graphqlQuery := `
		query Me {
			me {
				id
				username
				email
				firstName
				lastName
			}
		}
	`

	payload := map[string]interface{}{
		"query": graphqlQuery,
	}

	body, _ := json.Marshal(payload)
	// Dynamically resolve GraphQL URL
	graphqlURL := os.Getenv("RENDER_EXTERNAL_URL")
	if graphqlURL == "" {
		graphqlURL = "http://localhost:8080"
	}
	graphqlURL += "/query"

	// Make the request to your GraphQL API with the token
	req, err := http.NewRequest("POST", graphqlURL, bytes.NewReader(body))
	if err != nil {
		// Return empty object if there's an error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		// Return empty object if request fails or status is not OK
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
		return
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, _ := io.ReadAll(resp.Body)

	var graphqlResp models.MeResponse
	if err := json.Unmarshal(respBody, &graphqlResp); err != nil || graphqlResp.Data.Me == nil {
		// If there's an error in response or no user data, return empty object
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
		return
	}

	// Log the user in or set session if necessary
	// You can set a session here if using a session mechanism like JWT or server-side sessions.
	// Example: set a session here (implement according to your session management)
	// session := getSession(r)
	// session.Set("user", graphqlResp.Data.Me)

	// Respond with the user data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(graphqlResp.Data.Me)
}
