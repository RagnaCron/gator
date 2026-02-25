package main

import (
	"context"
	"os"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return err
	}

	os.Exit(0)
	return nil
}
