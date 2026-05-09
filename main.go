package main

import (
	"fmt"
	"net/http"

	"github.com/SarAvi21805/kdrama-tracker-backend/database"
	"github.com/SarAvi21805/kdrama-tracker-backend/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	database.InitDB()
	fmt.Println("Base de datos SQLite lista 🗄️")

	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)    // Muestra peticiones en la consola
	r.Use(middleware.Recoverer) // Evita que el servidor caiga ante error

	// CONFIGURACIÓN DE CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Cualquier origen (frontend)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// RUTAS
	r.Get("/series", handlers.GetSeries)
	r.Get("/series/{id}", handlers.GetSerieByID)
	r.Post("/series", handlers.CreateSerie)
	r.Put("/series/{id}", handlers.UpdateSerie)
	r.Delete("/series/{id}", handlers.DeleteSerie)

	fmt.Println("Servidor corriendo en http://localhost:8080 ✨🌸")
	http.ListenAndServe(":8080", r)
}
