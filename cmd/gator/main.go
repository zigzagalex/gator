package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq"
	"github.com/peterh/liner"
	"github.com/zigzagalex/gator/commands"
	"github.com/zigzagalex/gator/internal/config"
	"github.com/zigzagalex/gator/internal/database"
)

var (
	historyFile = filepath.Join(os.TempDir(), ".gator_history")
	names       = []string{"string1", "something2", "anything3", "nothing4"}
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

	// Setup liner
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	// Autocomplete setup
	line.SetCompleter(func(line string) (c []string) {
		for _, n := range names {
			if strings.HasPrefix(strings.ToLower(n), strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

	// Load liner history
	if f, err := os.Open(historyFile); err == nil {
		defer f.Close()
		line.ReadHistory(f)
	}
	// Set state
	state := &commands.State{
		DB:      dbQueries,
		Pointer: conf,
	}
	// Intitialize commantd
	cmdRegistry, err := commands.InitCommands()
	if err != nil {
		log.Fatalf("⚠️ Failed to initialize commands: %v", err)
	}

	fmt.Println("🐊 Gator — type 'exit' to quit")

	for {
		input, err := line.Prompt("> ")
		if err != nil {
			if err == liner.ErrPromptAborted {
				fmt.Println("\nAborted")
				break
			}
			log.Println("⚠️ Error reading line:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		line.AppendHistory(input)

		if input == "exit" {
			fmt.Println("Goodbye.")
			break
		}

		parts := strings.Fields(input)
		commandName := parts[0]
		commandArgs := parts[1:]

		cmd := commands.Command{
			Name: commandName,
			Args: commandArgs,
		}
		err = cmdRegistry.Run(state, cmd)
		if err != nil {
			fmt.Println("⚠️ Error:", err)
		}

	}
	// Save history
	if f, err := os.Create(historyFile); err == nil {
		defer f.Close()
		line.WriteHistory(f)
	}
}
