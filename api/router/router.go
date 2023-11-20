package router

import (
	"fmt"
	"strings"

	goaway "github.com/TwiN/go-away"
	"github.com/gin-gonic/gin"
	"github.com/lakrsv/profanity-crawler/api/swagger"
	"github.com/lakrsv/profanity-crawler/api/types"
	"github.com/lakrsv/profanity-crawler/internal/app/crawler"
	"github.com/thinkerou/favicon"
	"github.com/zc2638/swag"
)

type DefaultRouter struct {
}

func CreateRouter() *gin.Engine {
	router := gin.New()

	api := swagger.NewApi(DefaultRouter{})
	api.Walk(func(path string, e *swag.Endpoint) {
		h := e.Handler.(func(*gin.Context))
		path = swag.ColonPath(path)

		router.Handle(e.Method, path, h)
	})

	router.GET("/swagger/json", gin.WrapH(api.Handler()))
	router.GET("/swagger/ui/*any", gin.WrapH(swag.UIHandler("/swagger/ui", "/swagger/json", true)))

	router.SetTrustedProxies(nil)
	router.Use(favicon.New("./assets/favicon/favicon.ico"))

	return router
}

func (DefaultRouter) StartCrawl(c *gin.Context) {
	var crawlRequest types.StartCrawlRequest
	if err := c.BindJSON(&crawlRequest); err != nil {
		return
	}
	fmt.Println(crawlRequest)

	ch := make(chan crawler.CrawlResult)
	go crawler.Crawl(crawlRequest.Url, crawlRequest.Depth, ch)

	falsePositives := []string{
		"parse",
		"charset",
		"slutat",
		"thorny",
		"analogous",
		"assembly",
		"assemble",
		"ballstorp",
		"cockatoo",
		"sexual",
		"hearse",
		"hancock",
		"peacock",
		"assault",
		"button",
		"kick-ass",
		"starsector",
		"atwater",
		"smartwatch",
		"gamecocks",
		"stassi",
		"associée",
		"associant",
	}
	falsePositives = append(falsePositives, goaway.DefaultFalsePositives...)

	detector := goaway.NewProfanityDetector().WithSanitizeLeetSpeak(false).WithSanitizeSpecialCharacters(false).WithSanitizeAccents(false).WithSanitizeSpaces(false).WithCustomDictionary(goaway.DefaultProfanities, falsePositives, goaway.DefaultFalseNegatives)

	for res := range ch {
		if res.Error() != nil {
			fmt.Println("Got error", res.Error())
			continue
		}
		for _, line := range strings.Split(res.Body(), "\n") {
			swear := detector.ExtractProfanity(line)
			if swear != "" {
				if len(line) > 50 {
					line = line[:50] + "..."
				}
				fmt.Println("Swear found: ", swear, " in url: ", res.Path(), " in text: ", line)
			}
		}
	}
	fmt.Println("Done")
}

func (DefaultRouter) GetCrawl(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
}
