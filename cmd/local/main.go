package main

import (
	"github.com/lakrsv/profanity-crawler/api/router"
	"github.com/lakrsv/profanity-crawler/api/server"
)

func main() {
	server.ListenAndServe(8080, router.CreateRouter())
}
