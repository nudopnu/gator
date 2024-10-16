package main

import (
	"context"
	"fmt"

	"github.com/nudopnu/gator/internal/rss"
)

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
		fmt.Printf("%+v", item)
	}
	return nil
}
