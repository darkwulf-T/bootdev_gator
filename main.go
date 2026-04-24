package main

import (
	"fmt"

	"github.com/darkwulf-T/bootdev_gator/internal/config"
)

func main() {
	con, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = con.SetUser("darkwulf")
	if err != nil {
		fmt.Println(err)
		return
	}
	con, err = config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", con)
}
