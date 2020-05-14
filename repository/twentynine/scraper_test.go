package twentynine

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
			Url: "https://www.29cm.co.kr/product/663865",
			ExpectedTitle: "[스페셜 오더]_수피마 쿨맥스 티셔츠 16컬러",
			ExpectedBrand: "더니트컴퍼니",
			ExpectedImageSrc: "img.29cm.co.kr/next-product/2020/04/23/4a2540877c2a43a98ebd21edebdf4225_20200423183038.jpg?width=600",
		},
		{
			Url: "https://www.29cm.co.kr/product/673836",
			ExpectedTitle: "BALLOON PULLOVER (sky blue)_HJSXS20131SBL",
			ExpectedBrand: "하이드아웃",
			ExpectedImageSrc: "img.29cm.co.kr/next-product/2020/04/28/f3971f50fa214a27b752f45370276098_20200428154203.jpg?width=600",
		},
	}
	
	for _, scenario := range TestScenarios {
		item, err := Scrap(scenario.Url)
		assert.Nil(err)
		assert.Equal(scenario.ExpectedTitle, item.Title)
		assert.Equal(scenario.ExpectedBrand, item.Brand)
		assert.Equal(scenario.ExpectedImageSrc, item.ImageSrc)
	}
}