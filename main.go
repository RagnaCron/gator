package main

import (
	"log"
	"os"

	"github.com/ragnacron/gogator/internal/config"
)

func main() {
	c, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := &state{
		config: &c,
	}

	commands := commands{
		handlers: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)

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
