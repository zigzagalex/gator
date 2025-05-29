package commands

import "fmt"

func (c *Commands) Run(s *State, cmd Command) error {
	if handler, ok := c.Handlers[cmd.Name]; ok {
		return handler(s, cmd)
	}
	return fmt.Errorf("Unknown command: %s", cmd.Name)
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	if c.Handlers == nil {
		c.Handlers = make(map[string]func(*State, Command) error)
	}
	c.Handlers[name] = f
}

func InitCommands() (Commands, error) {
	cmdRegistry := Commands{
		Handlers: make(map[string]func(*State, Command) error),
	}

	cmdRegistry.Register("login", HandlerLogin)
	return cmdRegistry, nil
}
