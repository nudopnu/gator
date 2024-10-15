package main

import (
	"context"
	"fmt"

	"github.com/nudopnu/gator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	// if len(cmd.args) < 1 {
	// 	return errors.New("no url provided")
	// }
	// feedURL := cmd.args[0]
	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}
	fmt.Printf("%+v\n", feed)
	return nil
}
