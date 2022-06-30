package main

import (
	"fmt"
	"go-gorm-mux/api/config"
	"go-gorm-mux/api/controllers"
	"go-gorm-mux/api/database"
	"go-gorm-mux/api/middlewares"
	"go-gorm-mux/api/seed"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load the server config
	config := config.GetConfig()

	// Initialize database
	database.Connect(config)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register the routes
	RegisterRoutes(router)

	// Seed initial data for the database
	seed.Load(database.DB)

	// Start the server
	msg := fmt.Sprintf("Server running on port %s", config.Port)
	log.Printf("%s\n", msg)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Port), router))
}

// RegisterRoutes register all available routes for the API.
func RegisterRoutes(router *mux.Router) {
	// Home Route
	router.HandleFunc("/", middlewares.SetMiddlewareJSON(controllers.Home)).Methods("GET")

	// Login Route
	router.HandleFunc("/login", middlewares.SetMiddlewareJSON(controllers.Login)).Methods("POST")

	// Users routes
	router.HandleFunc("/users", middlewares.SetMiddlewareJSON(controllers.CreateUser)).Methods("POST")
	router.HandleFunc("/users", middlewares.SetMiddlewareJSON(controllers.GetUsers)).Methods("GET")
	router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(controllers.GetUser)).Methods("GET")
	router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(controllers.UpdateUser))).Methods("PUT")
	router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(controllers.DeleteUser)).Methods("DELETE")

	// Posts routes
	router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(controllers.CreatePost)).Methods("POST")
	router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(controllers.GetPosts)).Methods("GET")
	router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(controllers.GetPost)).Methods("GET")
	router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(controllers.UpdatePost))).Methods("PUT")
	router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(controllers.DeletePost)).Methods("DELETE")
}