package main

import (
	"context"
	"fmt"
	"time"

	"github.com/darkwulf-T/bootdev_gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", *feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("a name and url for the feed must be passed")
	}
	currentUserName := s.config.CurrentUserName
	currentUser, err := s.db.GetUser(context.Background(), currentUserName)
	if err != nil {
		return err
	}
	fname := cmd.arguments[0]
	furl := cmd.arguments[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      fname,
		Url:       furl,
		UserID:    currentUser.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("feed %q has been created\n", feed.Name)
	fmt.Printf("%+v\n", feed)
	return nil
}
