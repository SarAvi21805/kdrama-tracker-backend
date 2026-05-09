package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SarAvi21805/kdrama-tracker-backend/database"
	"github.com/SarAvi21805/kdrama-tracker-backend/models"
	"github.com/go-chi/chi/v5"
)

// Obtener series
func GetSeries(w http.ResponseWriter, r *http.Request) {
	// Consulta a la DB
	rows, err := database.DB.Query("SELECT id, title, genre, category, description, image_url, rating, is_favorite FROM series")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	series := []models.Serie{}
	for rows.Next() {
		var s models.Serie
		err := rows.Scan(&s.ID, &s.Title, &s.Genre, &s.Category, &s.Description, &s.ImageURL, &s.Rating, &s.IsFavorite)
		if err != nil {
			continue
		}
		series = append(series, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(series)
}

// Obtener una sola serie
func GetSerieByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var s models.Serie

	err := database.DB.QueryRow("SELECT id, title, genre, category, description, image_url, rating, is_favorite FROM series WHERE id = ?", id).
		Scan(&s.ID, &s.Title, &s.Genre, &s.Category, &s.Description, &s.ImageURL, &s.Rating, &s.IsFavorite)

	if err != nil {
		http.Error(w, "Serie no encontrada 🌸", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

// Eliminar serie
func DeleteSerie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id") // Obtiene el ID de la URL

	query := `DELETE FROM series WHERE id = ?`
	_, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Editar serie
func UpdateSerie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var s models.Serie
	json.NewDecoder(r.Body).Decode(&s)

	query := `UPDATE series SET title=?, genre=?, category=?, description=?, image_url=?, rating=?, is_favorite=? WHERE id=?`
	_, err := database.DB.Exec(query, s.Title, s.Genre, s.Category, s.Description, s.ImageURL, s.Rating, s.IsFavorite, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s)
}

// Recibe un JSON y lo inserta en la base de datos
func CreateSerie(w http.ResponseWriter, r *http.Request) {
	var s models.Serie
	// Decodifica el JSON que viene del Frontend
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Datos inválidos 🌸", http.StatusBadRequest)
		return
	}

	// Insertar en la BD
	query := `INSERT INTO series (title, genre, category, description, image_url, rating, is_favorite) 
              VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := database.DB.Exec(query, s.Title, s.Genre, s.Category, s.Description, s.ImageURL, s.Rating, s.IsFavorite)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Obtener ID generado y responder
	id, _ := result.LastInsertId()
	s.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(s)
}
