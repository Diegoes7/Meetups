package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Diegoes7/meetups/domain"
	"github.com/Diegoes7/meetups/graph"
	"github.com/Diegoes7/meetups/handlers"
	"github.com/Diegoes7/meetups/loader"
	customMiddleware "github.com/Diegoes7/meetups/middleware"
	"github.com/Diegoes7/meetups/models"
	"github.com/Diegoes7/meetups/postgres"
	"github.com/Diegoes7/meetups/internal/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

// const defaultPort = "8080"
const defaultPort = "8080" // Change port to avoid GraphQL playground landing

func main() {
	ctx := context.Background()
	connections, err := db.NewConnections(ctx)
	if err != nil {
		log.Fatalf("Could not initialize connections: %v", err)
	}
	DB := connections.Postgres
	defer DB.Close()
	DB.AddQueryHook(&postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	userRepo := postgres.UserRepo{DB: DB}
	meetupRepo := postgres.MeetupRepo{DB: DB}
	invitationRepo := postgres.InvitationRepo{DB: DB}
	messageRepo := postgres.MessageRepo{DB: DB}

	//$ Mux object is the main router that handles HTTP requests
	//! chi.NewRouter is used to instantiate a new Mux, which is the main router
	//! will route incoming HTTP requests to the appropriate handlers based on defined routes and middlewares.
	//& used to define routes, handle HTTP requests, and set up middleware.
	router := chi.NewRouter()

	//! Initialize CORS options
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*", "http:localhost:8080"}, //! Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"X-Custom-Header", "Content-Type"},
		ExposedHeaders:   []string{"Sec-WebSocket-Protocol"},
		AllowCredentials: true,
		Debug:            true,
	}

	// Create a new CORS handler
	corsHandler := cors.New(corsOptions).Handler

	//$ middlewares
	router.Use(corsHandler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(userRepo))

	d := domain.NewDomain(userRepo, meetupRepo, invitationRepo, messageRepo)

	resolver := &graph.Resolver{Domain: d}

	c := graph.Config{Resolvers: resolver}

	srv := handler.New(graph.NewExecutableSchema(c))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	}) //! realtime

	srv.Use(extension.Introspection{})

	queryHandler := srv

	// Initialize the SubscriptionManager
	graph.SubManager = graph.NewSubscriptionManager()

	//! Serve template-specific JS files
	router.Handle("/templates/static/*",
		http.StripPrefix("/templates/static/",
			http.FileServer(http.Dir("templates/static"))))

	//! Serve HTML on "/"
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.gohtml", "templates/meetups.gohtml", "templates/login.gohtml", "templates/create_meetup_modal.gohtml")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	router.Get("/meetup/{meetupID}", func(w http.ResponseWriter, r *http.Request) {
		meetupID := chi.URLParam(r, "meetupID")

		// Dereference the double pointer to access the actual domain object
		log.Printf("Fetching meetup with ID: %s from the database", meetupID)
		meetup, err := (*d).MeetupRepo.GetByID(meetupID)
		if err != nil {
			log.Printf("Error loading single meetup %s", err)
			http.Error(w, "Meetup not found", http.StatusNotFound)
			return
		}

		data := struct {
			Meetup *models.Meetup
		}{
			Meetup: meetup,
		}

		tmpl, err := template.ParseFiles("templates/meetup_chat.gohtml", "templates/messages.gohtml")
		if err != nil {
			log.Printf("Template parse error: %s", err)
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Printf("Template execution error: %s", err)
			http.Error(w, "No messages yet.", http.StatusInternalServerError)
		}
	})

	router.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/register.gohtml")
		if err != nil {
			http.Error(w, "Failed to load registration page", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	router.Get("/me", handlers.MeHandler)
	router.Post("/login", handlers.LoginHandler)
	router.Post("/logout", handlers.LogoutHandler)

	router.Post("/invite", handlers.InviteUserHandler(d))
	router.Get("/api/users", handlers.UsersHandler(d))
	router.Handle("/query", loader.DataLoaderMiddleware(DB, queryHandler))

	// router.Handle("/subscriptions", srv)

	// http.HandleFunc("/subscriptions", handlers.HandleSubscription)
	router.HandleFunc("/subscriptions", handlers.HandleSubscription)

	// router.Get("/meetup/{meetupID}", func(w http.ResponseWriter, r *http.Request) {
	// 	tmpl, err := template.ParseFiles("templates/meetup_chat.gohtml")
	// 	if err != nil {
	// 		http.Error(w, "Error loading template", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	tmpl.Execute(w, nil)
	// })

	//! Move GraphQL playground to "/playground"
	router.Handle("/playground", playground.Handler("GraphQL playground", "/query"))

	//! GraphQL API route
	router.Handle("/query", loader.DataLoaderMiddleware(DB, queryHandler))
	// http.Handle("/query", handler.GraphQL(graph.NewExecutableSchema(c)))

	log.Printf("connect to http://localhost:%s/ for GraphQL actual website", port)
	log.Printf("connect to http://localhost:%s/playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

