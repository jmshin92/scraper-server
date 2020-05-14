package musinsa

import (
	"context"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"github.com/jmshin92/scraper/item"
	"github.com/jmshin92/scraper/repository"
	"strings"
)

const (
	UrlPrefixMusinsa = "https://m.store.musinsa.com/"
)

const (
	titleSelector = `body.page-product-detail div.musinsa-wrapper div#product_order_info h2.prd-title`
	brandSelector = `body.page-product-detail div.musinsa-wrapper div#product_order_info div.wrap-product-stats ul.prd-state-table li.box-product-num div.stats a`
	imageSelector = `body.page-product-detail div.musinsa-wrapper div#product_order_info div.wrap-prd-thumb div#wrapper ul#thelist li.swiper-slide img`
)

func Scrap(url string) (*item.Item, error) {
	ctx, cancel := repository.NewContext()
	defer cancel()
	
	item := &item.Item{
		Url:url,
	}
	if err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(titleSelector),
		
		chromedp.Text(titleSelector, &item.Title),
		chromedp.Text(brandSelector, &item.Brand),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			img, err := dom.QuerySelector(node.NodeID, imageSelector).Do(ctx)
			if err != nil {
				return err
			}
			attrs, err := dom.GetAttributes(img).Do(ctx)
			if err != nil {
				return err
			}
			
			for i := 0; i < len(attrs); i += 2 {
				if attrs[i] == "src" {
					url := strings.TrimPrefix(attrs[i+1], "//")
					item.ImageSrc = url
					return nil
				}
			}
			return err
		}),
	); err != nil {
		return nil, err
	}
	return item, nil
}