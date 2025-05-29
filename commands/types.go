package commands

import "github.com/zigzagalex/gator/internal/config"

type State struct {
	Pointer *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Handlers map[string]func(*State, Command) error
}