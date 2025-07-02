package main

import (
	// loggerservice "github.com/avalokitasharma/lld/logger-service"
	// ratingservice "github.com/avalokitasharma/lld/rating-service"
	sqldb "github.com/avalokitasharma/lld/sql-db"
)

func main() {
	// loggerservice.Init()
	// ratingservice.Init()
	sqldb.Init()
}
