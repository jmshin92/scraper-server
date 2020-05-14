package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmshin92/scraper/controllers/api/scrap"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong",})
	})
	
	
	r.GET("/scrap", scrap.Get)
	
	
	return r
}