package commands

import "fmt"

func (c *Commands) Run(s *State, cmd Command) error {
	if handler, ok := c.Handlers[cmd.Name]; ok {
		return handler.Handler(s, cmd)

	}
	return fmt.Errorf("Unknown command: %s", cmd.Name)
}

func (c *Commands) CommandRegister(name, description, usage string, handler func(*State, Command) error) {
	if c.Handlers == nil {
		c.Handlers = make(map[string]CommandHandler)
	}
	c.Handlers[name] = CommandHandler{
		Description: description,
		Usage:       usage,
		Handler:     handler,
	}
}

func InitCommands() (Commands, error) {
	cmdRegistry := Commands{
		Handlers: make(map[string]CommandHandler),
	}

	cmdRegistry.CommandRegister("login", "Login as user", "login <username>", HandlerLogin)
	cmdRegistry.CommandRegister("register", "Register new user", "register <username>", HandlerRegister)
	cmdRegistry.CommandRegister("reset", "Reset databases, DEV only.", "reset", HandlerReset)
	cmdRegistry.CommandRegister("users", "List all users", "users", HandlerUsers)
	cmdRegistry.CommandRegister("agg", "Fetch and aggregate feeds", "agg <int><unit>", HandlerAgg)
	cmdRegistry.CommandRegister("addfeed", "Add RSS feed", "addfeed <feed_name> <url>", middlewareLoggedIn(HandlerAddFeed))
	cmdRegistry.CommandRegister("feeds", "List all feeds", "feeds", HandlerFeeds)
	cmdRegistry.CommandRegister("follow", "Follow a feed", "follow <feed_url>", middlewareLoggedIn(HandlerFollow))
	cmdRegistry.CommandRegister("following", "Show followed feeds by current user", "following", HandlerFollowing)
	cmdRegistry.CommandRegister("unfollow", "Unfollow a feed", "unfollow <feed_url>", middlewareLoggedIn(HandlerUnfollow))
	cmdRegistry.CommandRegister("b", "Browse posts", "b <int>", middlewareLoggedIn(HandlerBrowse))
	cmdRegistry.CommandRegister("help", "Show available commands", "help", HandlerHelp(&cmdRegistry))

	return cmdRegistry, nil
}
