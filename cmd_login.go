package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("login expects a single argument: <username>")
	}

	userName := cmd.args[0]
	err := s.config.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("login has set user: %s\n", userName)

	return nil
}
