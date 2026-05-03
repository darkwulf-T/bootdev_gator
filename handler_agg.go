package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/darkwulf-T/bootdev_gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("please select a time between requests")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("please select an actual time duration")
	}
	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			log.Printf("error scraping feeds: %v", err)
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return err
	}
	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}
	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
	}
	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("a name and url for the feed must be passed")
	}
	fname := cmd.arguments[0]
	furl := cmd.arguments[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      fname,
		Url:       furl,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("feed %q has been added\n", feed.Name)
	fmt.Printf("%+v\n", feed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	if len(feeds) == 0 {
		fmt.Println("No feeds found")
		return nil
	}
	for _, feed := range feeds {
		fmt.Printf("- Feed: %s, URL: %s, User: %s\n", feed.Name, feed.Url, feed.UserName)
	}
	return nil
}
