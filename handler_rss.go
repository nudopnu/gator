package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nudopnu/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("no duration provided")
	}
	duration := cmd.args[0]
	milliseconds, err := time.ParseDuration(duration)
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %s\n", duration)
	ticker := time.NewTicker(milliseconds)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func handlerAddFeed(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) < 2 {
		return errors.New("please provide a name and a url")
	}
	name := cmd.args[0]
	feedURL := cmd.args[1]
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       feedURL,
		UserID:    currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}
	fmt.Printf("Successfully added feed:\n%+v\n", feed)
	ff, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
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

func handlerFollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) < 1 {
		return errors.New("no url provided")
	}
	url := cmd.args[0]

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

func handlerFollowing(s *state, _ command, currentUser database.User) error {
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("error getting follows: %w", err)
	}
	fmt.Printf("%+v\n", follows)
	return nil
}

func handlerUnfollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.args) < 1 {
		return errors.New("no url to unfollow provided")
	}
	url := cmd.args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("there is no feed matching url '%s' %w", url, err)
	}
	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: currentUser.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error removing feedfollow: %w", err)
	}
	return nil
}
