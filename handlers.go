package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("login handler expects a username argument")
	}
	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Printf("user %s has been set\n", cmd.arguments[0])
	return nil
}
