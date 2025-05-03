package graph

import (
	"sync"

	"github.com/Diegoes7/meetups/models"
)

// SubscriptionManager handles subscriptions for meetup updates
type SubscriptionManager struct {
	mu          sync.Mutex
	subscribers map[string][]chan *models.MeetupUpdate
}

// NewSubscriptionManager creates a new instance of SubscriptionManager
func NewSubscriptionManager() *SubscriptionManager {
	return &SubscriptionManager{
		subscribers: make(map[string][]chan *models.MeetupUpdate),
	}
}

// Subscribe to a meetup's updates
func (s *SubscriptionManager) Subscribe(meetupID string) <-chan *models.MeetupUpdate {
	s.mu.Lock()
	defer s.mu.Unlock()

	ch := make(chan *models.MeetupUpdate, 1)
	s.subscribers[meetupID] = append(s.subscribers[meetupID], ch)
	return ch
}

// Publish updates to all subscribers of a meetup
func (s *SubscriptionManager) Publish(meetupID string, update *models.MeetupUpdate) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, ch := range s.subscribers[meetupID] {
		ch <- update
	}
}

// Global instance of SubscriptionManager
var SubManager = NewSubscriptionManager()
