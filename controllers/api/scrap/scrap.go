package scrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmshin92/scraper/item"
	"github.com/jmshin92/scraper/repository/common"
	"github.com/jmshin92/scraper/repository/musinsa"
	"github.com/jmshin92/scraper/repository/ohou"
	"github.com/jmshin92/scraper/repository/twentynine"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const QueryUrl = "url"

func Get(c * gin.Context) {
	url := c.Query(QueryUrl)
	if len(url) == 0 {
		err := fmt.Errorf("url Query is required")
		logrus.Error(err)
		JsonError(c, err)
		return
	}
	
	var res *item.Item
	var err error
	
	switch {
	case strings.HasPrefix(url, twentynine.UrlPrefixTwentyNine):
		res, err = twentynine.Scrap(url)
	case strings.HasPrefix(url, ohou.UrlPrefixOhouse):
		res, err = ohou.Scrap(url)
	case strings.HasPrefix(url, musinsa.UrlPrefixMusinsa):
		res, err = musinsa.Scrap(url)
	default:
		res, err = common.Scrap(url)
	}
	
	if err != nil {
		JsonError(c, err)
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func JsonError(c *gin.Context, err error) {
	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}