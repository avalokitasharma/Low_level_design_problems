package main

import (
	loggerservice "github.com/avalokitasharma/lld/logger-service"
	ratingservice "github.com/avalokitasharma/lld/rating-service"
)

func main() {
	loggerservice.Init()
	ratingservice.Init()
}
