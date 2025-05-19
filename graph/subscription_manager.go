package graph

import (
	"log"
	"sync"

	"github.com/Diegoes7/meetups/models"
)

// SubscriptionManager handles subscriptions for meetup updates
type SubscriptionManager struct {
	mu          sync.Mutex
	subscribers map[string][]chan *models.MeetupUpdate
	active      map[string]bool // Track active status
}

// NewSubscriptionManager creates a new instance of SubscriptionManager
func NewSubscriptionManager() *SubscriptionManager {
	return &SubscriptionManager{
		subscribers: make(map[string][]chan *models.MeetupUpdate),
		active:      make(map[string]bool),
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

// Unsubscribe removes a specific channel from a meetup's subscribers and closes it
func (s *SubscriptionManager) Unsubscribe(meetupID string, ch <-chan *models.MeetupUpdate) {
	s.mu.Lock()
	defer s.mu.Unlock()

	channels := s.subscribers[meetupID]
	for i, subscriber := range channels {
		if subscriber == ch {
			// Close the channel and remove it from the slice
			close(subscriber)
			s.subscribers[meetupID] = append(channels[:i], channels[i+1:]...)
			break
		}
	}
}

func (s *SubscriptionManager) CloseMeetup(meetupID string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if subs, ok := s.subscribers[meetupID]; ok {
		for _, ch := range subs {
			close(ch) // Closing channel tells handler to stop
		}
		delete(s.subscribers, meetupID)
	}

	delete(s.active, meetupID)
}

func (s *SubscriptionManager) SetActive(meetupID string, active bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.active[meetupID] = active
	log.Printf("[SetActive] %s -> %v", meetupID, active)
}

func (s *SubscriptionManager) IsActive(meetupID string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	active := s.active[meetupID]
	log.Printf("[IsActive] %s -> %v", meetupID, active)
	return active
}

// func (s *SubscriptionManager) SetActive(meetupID string, active bool) {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	s.active[meetupID] = active
// }

// func (s *SubscriptionManager) IsActive(meetupID string) bool {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()
// 	return s.active[meetupID]
// }

// Global instance of SubscriptionManager
var SubManager = NewSubscriptionManager()
