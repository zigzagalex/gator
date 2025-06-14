package commands

import (
	"github.com/zigzagalex/gator/internal/config"
	"github.com/zigzagalex/gator/internal/database"
)

type State struct {
	DB      *database.Queries
	Pointer *config.Config
}

type Command struct {
	Name string
	Args []string
}

type CommandHandler struct {
	Description string
	Usage       string
	Handler     func(*State, Command) error
}

type Commands struct {
	Handlers map[string]CommandHandler
}
