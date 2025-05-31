package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/zigzagalex/gator/commands"
	"github.com/zigzagalex/gator/internal/config"
	"github.com/zigzagalex/gator/internal/database"
)

func main() {
	conf, _ := config.Read()

	db, err := sql.Open("postgres", conf.DBURL)
	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("DB ping failed: %v", err)
	}

	dbQueries := database.New(db)

	state := &commands.State{
		DB:      dbQueries,
		Pointer: conf,
	}

	cmdRegistry, _ := commands.InitCommands()

	if len(os.Args) < 2 {
		fmt.Println("No command provided.")
		os.Exit(1)
	}
	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	cmd := commands.Command{
		Name: commandName,
		Args: commandArgs,
	}

	err1 := cmdRegistry.Run(state, cmd)
	if err1 != nil {
		fmt.Println("Error:", err1)
		os.Exit(1)
	}

}
