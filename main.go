package main

import (
	"fmt"

	c "github.com/ragnacron/gogator/internal/config"
)

func main() {
	config := c.Read()
	config.SetUser("ragnacron")
	config = c.Read()
	fmt.Println(config.DBUrl)
	fmt.Println(config.CurrentUserName)
}
