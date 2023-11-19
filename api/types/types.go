package types

import "github.com/gin-gonic/gin"

type CrawlApi interface {
	StartCrawl(*gin.Context)
	GetCrawl(*gin.Context)
}

type StartCrawlRequest struct {
	Url   string `json:"url" required:"true"`
	Depth int    `json:"depth" required:"false"`
}

type StartCrawlResponse struct {
	Id string `json:"id"`
}

type GetCrawlResponse struct {
	Id string `json:"id"`
}
