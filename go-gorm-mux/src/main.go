package main

import (
	"fmt"
	"go-gorm-mux/src/api/config"
	"go-gorm-mux/src/api/controllers"
	"go-gorm-mux/src/api/database"
	"go-gorm-mux/src/api/middlewares"
	"go-gorm-mux/src/api/seed"
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
	log.Printf(fmt.Sprintf("Starting Server on port %s\n", config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Port), router))
}

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