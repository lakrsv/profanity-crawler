package swagger

import (
	"net/http"

	"github.com/lakrsv/profanity-crawler/api/types"
	"github.com/zc2638/swag"
	"github.com/zc2638/swag/endpoint"
	"github.com/zc2638/swag/option"
)

func NewApi(crawlApi types.CrawlApi) *swag.API {
	api := swag.New(
		option.Title("Profanity Crawler API Doc"),
		option.Endpoints(
			endpoint.New(
				http.MethodPost, "/crawl",
				endpoint.Handler(crawlApi.StartCrawl),
				endpoint.Summary("Start a crawl"),
				endpoint.Description("Crawls a given url to supplied depth"),
				endpoint.Body(types.StartCrawlRequest{}, "Crawl request", true),
				endpoint.Response(http.StatusOK, "Successfully started a crawl", endpoint.Schema(types.StartCrawlResponse{})),
			),
			endpoint.New(
				http.MethodGet, "/crawl/{id}",
				endpoint.Handler(crawlApi.GetCrawl),
				endpoint.Summary("Get a crawl by id"),
				endpoint.Path("id", "integer", "ID of crawl to return", true),
				endpoint.Response(http.StatusOK, "Successfully retrieved crawl", endpoint.Schema(types.GetCrawlResponse{})),
			),
		),
	)
	return api
}
