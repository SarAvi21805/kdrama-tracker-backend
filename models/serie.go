package models

// Estructura de un K-Drama o Anime
type Serie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	Category    string `json:"category"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Rating      int    `json:"rating"`
	IsFavorite  bool   `json:"is_favorite"`
}
