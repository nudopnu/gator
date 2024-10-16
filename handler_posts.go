package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/nudopnu/gator/internal/database"
)

func handlerBrowse(s *state, cmd command) error {
	limit := 2
	if len(cmd.args) > 0 {
		arg, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("error parsing cmd arg '%s': %w", cmd.args[0], err)
		}
		limit = arg
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		Name:  s.cfg.CurrentUserName,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error fetching posts: %w", err)
	}
	for _, post := range posts {
		fmt.Printf("***%s***\n", post.Title.String)
		fmt.Printf("%s\n", post.Description.String)
	}
	return nil
}
