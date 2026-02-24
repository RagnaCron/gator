package main

import (
	"errors"
	"log"
	"os"

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

	if len(os.Args) <= 2 {
		logFatalln(
			errors.New("missing arguments"),
		)
	}

	command := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	commands.run(&s, command)
}

func logFatalln(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
