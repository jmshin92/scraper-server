package common

import (
	"github.com/jmshin92/scraper/item"
	"github.com/julianshen/og"
)

func Scrap(url string) (*item.Item, error) {
	pageInfo, err := og.GetPageInfoFromUrl(url)
	if err != nil {
		return nil, err
	}
	
	res := &item.Item{
		Url: url,
	}
	
	res.Title = pageInfo.Title
	res.ImageSrc = pageInfo.Images[0].Url
	return res, nil
}