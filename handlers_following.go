package main

import (
	"context"
	"fmt"
	"time"

	"github.com/darkwulf-T/bootdev_gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("handler follow expects a url as argument")
	}
	url := cmd.arguments[0]
	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("no registered feed with this URL was found")
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Current user: %s\n", feedFollow.UserName)
	fmt.Printf("Followed feed: %s\n", feedFollow.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	if len(feeds) == 0 {
		fmt.Println("You aren't following any feeds yet.")
		return nil
	}
	fmt.Printf("Feeds followed by %s:\n", s.config.CurrentUserName)
	for _, feed := range feeds {
		fmt.Printf("   - %s\n", feed.FeedName)
	}
	return nil
}
