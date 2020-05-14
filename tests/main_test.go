package test

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"testing"
	
	"github.com/stretchr/testify/assert"
)


func TestScraper(t *testing.T) {
	assert := assert.New(t)
	
	
	testPing(assert)
}

func testPing(assert *assert.Assertions) {
	u := getLoopBackUrl("ping")
	logrus.Info(u.String())
	resp, err := http.Get(u.String())
	assert.Nil(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}

func getLoopBackUrl(path string) *url.URL {
	return &url.URL{
		Scheme: "http",
		Host: "localhost:9090",
		Path: path,
	}
}