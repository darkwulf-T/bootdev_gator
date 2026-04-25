package main

func newCommands() *commands {
	c := &commands{
		handlers: make(map[string]func(*state, command) error),
	}
	c.register("login", handlerLogin)
	return c
}
