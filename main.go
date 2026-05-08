package main

import (
	"fmt"
	"net/http"

	"github.com/SarAvi21805/kdrama-tracker-backend/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
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

	fmt.Println("Servidor mágico corriendo en http://localhost:8080 ✨🌸")
	http.ListenAndServe(":8080", r)
}
