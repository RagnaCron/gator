package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ragnacron/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := int32(2)
	if len(cmd.args) == 1 {
		if v, err := strconv.ParseInt(cmd.args[0], 10, 32); err == nil {
			limit = int32(v)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: limit,
	})
	if err != nil {
		return fmt.Errorf("couldn't get user posts: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("* Title:    %s\n", post.Title)
		fmt.Printf("%s\n", post.Description)
		fmt.Println("========================================")
	}

	return nil
}
