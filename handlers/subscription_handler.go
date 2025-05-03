package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Diegoes7/meetups/graph"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // In production, restrict origins
	},
}

// HandleSubscription handles WebSocket connections for subscriptions
func HandleSubscription(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "WebSocket upgrade failed", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	meetupID := r.URL.Query().Get("meetupID")
	if meetupID == "" {
		http.Error(w, "Missing meetupID parameter", http.StatusBadRequest)
		return
	}

	// Subscribe using SubscriptionManager
	ch := graph.SubManager.Subscribe(meetupID)
	log.Printf("Client subscribed to meetupID: %s", meetupID)

	// Optionally send welcome message
	_ = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Subscribed to %s", meetupID)))

	for update := range ch {
		err := conn.WriteJSON(update)
		if err != nil {
			log.Printf("WebSocket write error: %v", err)
			break // Exit loop if sending fails
		}
	}
}
