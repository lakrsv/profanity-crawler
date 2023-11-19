package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lakrsv/profanity-crawler/api/swagger"
	"github.com/lakrsv/profanity-crawler/api/types"
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
}

func (DefaultRouter) GetCrawl(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
}
