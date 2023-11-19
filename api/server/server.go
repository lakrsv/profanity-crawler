package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lakrsv/profanity-crawler/api/router"
)

func ListenAndServe(port int, r *gin.Engine) {
	router := router.CreateRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
