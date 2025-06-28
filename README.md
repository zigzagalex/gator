# gator
A CLI/TUI RSS feed aggregator using Go.

This is a personal project to learn:
1. How to integrate a Go application with a PostgeSQL database.
2. How to build a long-running service that continuously fetches new posts from RSS feeds and stores them in the database.
3. How to build a TUI (text user interface) using ([Bubbletea](https://github.com/charmbracelet/bubbletea/tree/main))
4. How to write a decent installation and user guide


## The why and the what
This is a simple REPL CLI/TUI that allows for users to follow and fetch RSS feeds. All in the terminal, because honestly is there a better UI than a terminal?

The posts can then be opened to be read in the default browser. 

VIDEO

## How to install
1. Make sure you have PostgreSQL (>=15.0) and Go installed.

2. Install CLI using ```go install github.com/zigzagalex/gator```

3. Create a config file for the database using ```touch gator/internal/config/.gatorconfig.json``` and write 
```
{
    "db_url": "postgres://user:password@localhost:port/db_name?sslmode=disable",
    "current_user_name": "example_user"
}
```
and set the db_url to your database url, so that gator can connect to it. You can also set the session user name with current_user_name. 

4. Run migrations, by running ```goose -dir ./gator/sql/schema postgres "<db_url>" up``` where you replace the db_url with your actual link. 

5. Run gator ```./gator-tui``` the commands are explained on the bottom. 


Alternative for the brave: if you trust my executable script and have a mac running MacOS run ```chmod +x scripts/setup_db.sh``` out of the root of the project and then run ```./scripts/setup_db.sh ```. It should do steps 1-4 for you. 

## Commands
* Arrows for up, down, next/prev page
* (enter) for selection
* (esc) for going back/closing the CLI/TUI
* (+) for adding new user/feed
* (-) for deleting user/unfollowing feed
* (=) for following feed stored in the database

## What could be improved?
- Add an HTTP API (and authentication/authorization) that allows other users to interact with the service remotely
- Use docker for installation, to make the setup experience better

