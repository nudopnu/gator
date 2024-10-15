package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nudopnu/gator/internal/database"
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

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("please provide a name and a url")
	}
	name := cmd.args[0]
	feedURL := cmd.args[1]
	currentUser := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("invalid current user '%s': %w", currentUser, err)
	}
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       feedURL,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}
	fmt.Printf("Successfully added feed:\n%+v\n", feed)
	ff, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error following: %w", err)
	}
	fmt.Printf("Successfully followed:\n%+v\n", ff)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feeds: %w", err)
	}
	fmt.Printf("%+v\n", feeds)
	return nil
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("no url provided")
	}
	url := cmd.args[0]

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("invalid current user '%s': %w", currentUser, err)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error fetching FeedFollow: %w", err)
	}
	ff, err := s.db.CreateFeedFollow(context.TODO(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating FeedFollow: %w", err)
	}
	fmt.Printf("Successfully followed.\n%+v\n", ff)
	return nil
}

func handlerFollowing(s *state, _ command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("invalid current user '%s': %w", s.cfg.CurrentUserName, err)
	}
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("error getting follows: %w", err)
	}
	fmt.Printf("%+v\n", follows)
	return nil
}
