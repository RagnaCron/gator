package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ragnacron/gator/internal/database"
)

const timeLayout = "Wed 11 Feb 2026 00:00:00 +0000"

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time_duration>", cmd.name)
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s...\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		if err = scrapFeeds(s); err != nil {
			fmt.Println(err)
		}
	}
}

func scrapFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get next feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:            nextFeed.ID,
		UpdatedAt:     time.Now().UTC(),
		LastFetchedAt: sql.NullTime{Valid: true, Time: time.Now().UTC()},
	})

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch next feed: %w", err)
	}

	fmt.Printf("Fetched feed %s\n", feed.Channel.Title)

	for _, item := range feed.Channel.Item {
		parsedTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return err
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: parsedTime,
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			return fmt.Errorf("couldn't save post: %w", err)
		}
	}
	fmt.Println("Successfully saved posts!")

	return nil
}
