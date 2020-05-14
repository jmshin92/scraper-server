package ohou

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
			Url: "https://ohouse.onelink.me/2107755860/c9c4abab",
			ExpectedTitle: "[오늘의딜][TV방영특가] 나 혼자 스윙! 스윙 빨래바구니 2단/3단 모음",
			ExpectedBrand: "네이쳐리빙",
			ExpectedImageSrc: "https://image.ohou.se/i/bucketplace-v2-development/uploads/productions/1560995346651_Lp6290ko.jpg?gif=1&w=480&h=480&c=c",
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