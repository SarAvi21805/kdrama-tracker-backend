package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SarAvi21805/kdrama-tracker-backend/models"
)

// Base de datos temporal
var series = []models.Serie{
	{
		ID:          1,
		Title:       "Guardian: The Lonely God 🌸",
		Genre:       "Fantasía / Romance",
		Category:    "K-Drama",
		Description: "Un ser inmortal busca a su novia humana.",
		ImageURL:    "https://is3-ssl.mzstatic.com/image/thumb/EnV6baw9ITh_lCEHrhACfg/1200x675.jpg",
		Rating:      5,
		IsFavorite:  false,
	},
	{
		ID:          2,
		Title:       "Demon Slayer",
		Genre:       "Acción",
		Category:    "Anime",
		Description: "Tanjiro busca curar a su hermana convertida en demonio.",
		ImageURL:    "https://tse4.mm.bing.net/th/id/OIP.rGPe8LN9nbqGwBlmOkDLbQHaEK?w=1200&h=675&rs=1&pid=ImgDetMain&o=7&rm=3",
		Rating:      5,
		IsFavorite:  false,
	},
	{
		ID:          3,
		Title:       "Spy x Family",
		Genre:       "Comedia/Acción",
		Category:    "Anime",
		Description: "Un espía, una asesina y una telépata forman una familia falsa.",
		ImageURL:    "https://tse1.mm.bing.net/th/id/OIP.2ZPe0nk9BPmWXdVDvAg8DQHaD5?w=1000&h=526&rs=1&pid=ImgDetMain&o=7&rm=3",
		Rating:      5,
		IsFavorite:  false,
	},
}

// Responde con la lista de las series en formato JSON
func GetSeries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(series)
}
