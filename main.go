package main

import (
	"fmt"
	"log"

	c "github.com/ragnacron/gogator/internal/config"
)

func main() {
	config, err := c.Read()
	logFatalln(err)

	err = config.SetUser("ragnacron")
	logFatalln(err)

	config, err = c.Read()
	logFatalln(err)

	fmt.Println(config.DBUrl)
	fmt.Println(config.CurrentUserName)
}

func logFatalln(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
