package ohou

import (
	"github.com/chromedp/chromedp"
	"github.com/jmshin92/scraper/item"
	"github.com/jmshin92/scraper/repository"
)

const (
	UrlPrefixOhouse = "https://ohouse.onelink.me/"
)

const (
	overviewSelector = `body div.layout div.production-selling div.production-selling-overview div.production-selling-overview__container `

	contentSelector = overviewSelector + `div.production-selling-overview__content div.production-selling-header `
	titleSelector = contentSelector + `h1.production-selling-header__title span.production-selling-header__title__name`
	brandSelector = contentSelector + `h1.production-selling-header__title p.production-selling-header__title__brand-wrap`

	imageSelector = overviewSelector + `div.production-selling-cover-image-container img.production-selling-cover-image__entry__image`
)

func Scrap(url string) (*item.Item, error) {
	ctx, cancel := repository.NewContext()
	defer cancel()
	
	item := &item.Item{
		Url: url,
	}
	if err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(titleSelector),

		chromedp.Text(titleSelector, &item.Title),
		chromedp.Text(brandSelector, &item.Brand),
		chromedp.AttributeValue(imageSelector, "src", &item.ImageSrc, nil),
	); err != nil {
		return nil, err
	}
	
	return item, nil
}