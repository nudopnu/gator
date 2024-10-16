package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nudopnu/gator/internal/database"
	"github.com/nudopnu/gator/internal/rss"
)

func parseDate(date string) (time.Time, error) {
	layouts := []string{
		"Mon, 2 Jan 2006 15:04:05 -0700",
		"Mon, 02 Jan 2006 15:04:05 -0700",
		"2 Jan 2006 15:04:05 -0700",
	}

	var parsedDate time.Time
	var err error

	for _, layout := range layouts {
		parsedDate, err = time.Parse(layout, date)
		if err == nil {
			return parsedDate, nil
		}
	}
	return time.Time{}, fmt.Errorf("failed to parse date '%s': %w", date, err)
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching next feed: %w", err)
	}
	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("error marking %+v as fetched: %w", feed, err)
	}
	rss, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("error fetching rss feed ''%s': %w", feed.Url, err)
	}
	for _, item := range rss.Channel.Item {
		publishedAt, err := parseDate(item.PubDate)
		if err != nil {
			fmt.Println(err)
			continue
		}
		s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       sql.NullString{String: item.Title, Valid: true},
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
	}
	fmt.Printf("Fetched posts from feed '%s'\n", feed.Name)
	return nil
}
