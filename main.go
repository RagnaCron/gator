package main

import (
	"log"

	"github.com/ragnacron/gogator/internal/config"
)

func main() {
	c, err := config.Read()
	logFatalln(err)

	s := state{
		config: &c,
	}

	commands := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
}

func logFatalln(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
