package commands

import "fmt"

func (c *Commands) Run(s *State, cmd Command) error {
	if handler, ok := c.Handlers[cmd.Name]; ok {
		return handler(s, cmd)
	}
	return fmt.Errorf("Unknown command: %s", cmd.Name)
}

func (c *Commands) CommandRegister(name string, f func(*State, Command) error) {
	if c.Handlers == nil {
		c.Handlers = make(map[string]func(*State, Command) error)
	}
	c.Handlers[name] = f
}

func InitCommands() (Commands, error) {
	cmdRegistry := Commands{
		Handlers: make(map[string]func(*State, Command) error),
	}

	cmdRegistry.CommandRegister("login", HandlerLogin)
	cmdRegistry.CommandRegister("register", HandlerRegister)
	cmdRegistry.CommandRegister("reset", HandlerReset)
	cmdRegistry.CommandRegister("users", HandlerUsers)
	cmdRegistry.CommandRegister("agg", HandlerAgg)
	cmdRegistry.CommandRegister("addfeed", middlewareLoggedIn(HandlerAddFeed))
	cmdRegistry.CommandRegister("feeds", HandlerFeeds)
	cmdRegistry.CommandRegister("follow", middlewareLoggedIn(HandlerFollow))
	cmdRegistry.CommandRegister("following", HandlerFollowing)
	cmdRegistry.CommandRegister("unfollow", middlewareLoggedIn(HandlerUnfollow))
	cmdRegistry.CommandRegister("b", middlewareLoggedIn(HandlerBrowse))
	cmdRegistry.CommandRegister("help", HandlerHelp(&cmdRegistry))

	return cmdRegistry, nil
}
