# Gator

![Repo size](https://img.shields.io/github/repo-size/darkwulf-T/bootdev_gator)
![Last commit](https://img.shields.io/github/last-commit/darkwulf-T/bootdev_gator)

## Project Description
This project is a CLI tool, created as part of the Boot.Dev Back-End Developer Path as a guided project. It is a Blog aggregator that allowing you to add RSS feeds and track changes to those feeds.

---

## Installation

### Prerequisites
- go 1.26.1+
- Postgres 16.13+

### Installing the tool
**Installation**

Run the following command in your console:
```bash
go install github.com/darkwulf-T/bootdev_gator
```

**Set up config file** 

In your home directory create a file called `.gatorconfig.json` with the following content:
```
{
    "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable"
}
```
Replace `username`, `password` and `gator`with your actual Postgres credentials and database name.

**Set up Database**

Create the following tables in Postgres using `psql`:
```sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT UNIQUE NOT NULL
);

CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    last_fetched_at TIMESTAMP
);

CREATE TABLE feed_follows (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE(user_id, feed_id)
);

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    description TEXT,
    published_at TIMESTAMP,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);
```

---
## Usage 
### Command list:
- login:
    - Usage: `gator login <username>`
    - Log in as the selected user (if they exist)
- register:
    - Usage: `gator register <username>`
    - Create and log in with the selected username
- users:
    - Usage:  `gator users`
    - Prints all registered users to the terminal
- agg:
    - Usage: `gator agg <time-between-requests>`
    -  Starts the feed aggregator. Continuously fetches posts from all registered feeds on  a timer until stopped with Ctrl+C. The `time-between_requests`arguments controls how often feeds are fetched (e.g. `1m`, `5m`, ...). Please select a reasonable timeframe to not send too many requests.
- addfeed:
    - Usage: `gator addfeed <name> <url>`
    - Adds a feed to the database. The user adding a feed will automatically follow the feed.
- feeds: 
    - Usage: `gator feeds`
    - Prints all feeds in the database to the terminal.
- follow: 
    - Usage: `gator follow <url>`
    - Allows a user to follow a feed created by another user.
- following:
    - Usage: `gator following`
    - Prints all feeds the current user is following to the terminal.
- unfollow: 
    - Usage: `gator unfollow <url>`
    - Let the user unfollow a selected feed.
- browse: 
    - Usage: `gator browse <limit>`
    - Shows the newest posts from all feeds the user follows. The `limit`argument is an optional argument which controls how many posts are displayed. Its default value is 2.

### How to use the tool:
After setting everything up (see above) use the `register` command to create a user and the `addfeed` command to add some RSSFeeds. Use the `agg` command in a background terminal to start fetching posts. With the `browse` command the newest posts can be displayed.