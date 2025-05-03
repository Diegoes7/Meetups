// package handlers

// import (
// 	"html/template"
// 	"log"
// 	"net/http"

// 	"github.com/Diegoes7/meetups/models"
// 	"github.com/go-chi/chi/v5"
// )

// func Meetup(w http.ResponseWriter, r *http.Request) {
// 	meetupID := chi.URLParam(r, "meetupID")

// 	// Dereference the double pointer to access the actual domain object
// 	log.Printf("Fetching meetup with ID: %s from the database", meetupID)
// 	meetup, err := (*d).MeetupRepo.GetByID(meetupID)
// 	if err != nil {
// 		log.Printf("Error loading single meetup %s", err)
// 		http.Error(w, "Meetup not found", http.StatusNotFound)
// 		return
// 	}

// 	data := struct {
// 		Meetup *models.Meetup
// 	}{
// 		Meetup: meetup,
// 	}

// 	tmpl, err := template.ParseFiles("templates/meetup_chat.gohtml")
// 	if err != nil {
// 		log.Printf("Template parse error: %s", err)
// 		http.Error(w, "Error loading template", http.StatusInternalServerError)
// 		return
// 	}

// 	err = tmpl.Execute(w, data)
// 	if err != nil {
// 		log.Printf("Template execution error: %s", err)
// 		http.Error(w, "Error rendering page", http.StatusInternalServerError)
// 	}
// 	// log.Printf("Received request for meetup with ID: %s", chi.URLParam(r, "meetupID"))
// 	// handlers.Meetup(w, r, &d)
// }

package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Diegoes7/meetups/domain"
	"github.com/Diegoes7/meetups/models"
	"github.com/go-chi/chi/v5"
)

// This returns a handler function and closes over the domain
func MeetupHandler(d *domain.Domain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		meetupID := chi.URLParam(r, "meetupID")
		log.Printf("Fetching meetup with ID: %s from the database", meetupID)

		if meetupID == "" {
			http.Error(w, "Missing meetup ID", http.StatusBadRequest)
			return
		}

		meetup, err := d.MeetupRepo.GetByID(meetupID)
		if err != nil {
			log.Printf("Error loading single meetup: %s", err)
			http.Error(w, "Meetup not found", http.StatusNotFound)
			return
		}

		data := struct {
			Meetup *models.Meetup
		}{
			Meetup: meetup,
		}

		tmpl, err := template.ParseFiles("templates/meetup_chat.gohtml")
		if err != nil {
			log.Printf("Template parse error: %s", err)
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Printf("Template execution error: %s", err)
			http.Error(w, "Error rendering page", http.StatusInternalServerError)
		}
	}
}

