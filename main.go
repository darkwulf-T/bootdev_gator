package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/darkwulf-T/bootdev_gator/internal/config"
	"github.com/darkwulf-T/bootdev_gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	con, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	currentState := state{config: &con}

	dbURL := con.DbUrl
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	currentState.db = dbQueries

	handlers := newCommands()
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "please select a command to execute")
		os.Exit(1)
	}
	cmd := command{
		name:      args[0],
		arguments: args[1:],
	}
	err = handlers.run(&currentState, cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
