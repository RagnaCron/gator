package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ragnacron/gogator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Printf("Feed: %+v\n", feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	user, err := s.db.GetUser(s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	})
	return nil
}
