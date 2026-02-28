package main

import (
	"strconv"

	"github.com/ragnacron/gator/internal/database"
)

func handler_browse(s *state, cmd command, user database.User) error {
	limit := int32(2)
	if len(cmd.args) == 1 {
		if v, err := strconv.ParseInt(cmd.args[0], 10, 32); err == nil {
			limit = int32(v)
		}
	}
	return nil
}
