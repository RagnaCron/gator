package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/ragnacron/gogator/internal/config"
	"github.com/ragnacron/gogator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func main() {
	c, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := &state{
		config: &c,
	}

	db, err := sql.Open("postgres", c.DBUrl)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	s.db = database.New(db)

	commands := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerListUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerListFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))

	if len(os.Args) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
	}

	cName := os.Args[1]
	cArgs := os.Args[2:]

	err = commands.run(s, command{name: cName, args: cArgs})
	if err != nil {
		log.Fatal(err)
	}
}
