package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Para el driver
)

var DB *sql.DB

func InitDB() {
	var err error
	// Abre la conexión.
	DB, err = sql.Open("sqlite", "./series.db")
	if err != nil {
		log.Fatal("Error al abrir la base de datos: ", err)
	}

	// Crear la tabla si no existe
	query := `
	CREATE TABLE IF NOT EXISTS series (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		genre TEXT,
		category TEXT,
		description TEXT,
		image_url TEXT,
		rating INTEGER,
		is_favorite BOOLEAN DEFAULT 0
	);`

	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal("Error al crear la tabla: ", err)
	}
}
