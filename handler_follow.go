package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ragnacron/gogator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Println("Feed follow created:")
	printFollow(follow)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find feed follows: %w", err)
	}

	if len(follows) == 0 {
		return fmt.Errorf("No feed follows found for this user.")
	}

	fmt.Printf("Found %d feed follows for %s:\n", len(follows), s.config.CurrentUserName)
	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}

	return nil
}

func printFollow(f database.CreateFeedFollowRow) {
	fmt.Printf("* User name:    %s\n", f.UserName)
	fmt.Printf("* Feed name:    %s\n", f.FeedName)
}
