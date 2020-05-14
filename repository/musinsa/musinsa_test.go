package musinsa

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestScraper(t *testing.T) {
	assert := assert.New(t)
	
	TestScenarios := []struct {
		Url string
		ExpectedTitle string
		ExpectedBrand string
		ExpectedImageSrc string
	} {
		{
			Url: "https://m.store.musinsa.com/app/product/detail/1390339/0",
			ExpectedTitle: "[패키지] 9TH ANNIVERSARY 3PACK T-SHIRTS EDITION",
			ExpectedBrand: "GROOVE RHYME",
			ExpectedImageSrc: "image.msscdn.net/images/goods_img/20200408/1390339/1390339_1_500.jpg",
		},
		{
			Url: "https://m.store.musinsa.com/app/product/detail/1442667/0",
			ExpectedTitle: "아드리안 시스루 숏 봄버 awa285w(PALE JADE)",
			ExpectedBrand: "ANDERSSON BELL for WOMEN",
			ExpectedImageSrc:"image.msscdn.net/images/goods_img/20200512/1442667/1442667_1_500.jpg",
		},
	}
	
	for _, scenario := range TestScenarios {
		item ,err := Scrap(scenario.Url)
		assert.Nil(err)
		assert.Equal(scenario.ExpectedTitle, item.Title)
		assert.Equal(scenario.ExpectedBrand, item.Brand)
		assert.Equal(scenario.ExpectedImageSrc, item.ImageSrc)
	}
}