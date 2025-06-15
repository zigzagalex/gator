package main

import (
	"database/sql"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/lib/pq"
	"github.com/zigzagalex/gator/internal/config"
	"github.com/zigzagalex/gator/internal/database"
	"github.com/zigzagalex/gator/internal/ui"
)

func main() {
	// Setup database connection
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("⚠️ Failed to read config: %v", err)
	}

	db, err := sql.Open("postgres", conf.DBURL)
	if err != nil {
		log.Fatalf("⚠️ Failed to connect to db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("⚠️ DB ping failed: %v", err)
	}
	// Get queries
	dbQueries := database.New(db)

	// Get users

	// Start UI
	m := NewUI(dbQueries)
	p := tea.NewProgram(m)
	_, err = p.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func NewUI(q *database.Queries) ui.Model {
	return ui.Model{
		Q:       q,
		Loading: true,
	}
}
