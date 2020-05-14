package twentynine

import (
	"context"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"github.com/jmshin92/scraper/item"
	"github.com/jmshin92/scraper/repository"
	"strings"
)

const (
	UrlPrefixTwentyNine = "https://www.29cm.co.kr/"
)

const (
	overviewSelector = `body div section.product div.detail_wrap div.detail_cnt_wrap div.detail_item `
	titleSelector = overviewSelector + `div.prd_info div.info div.name`
	brandSelector = overviewSelector + `div.item_detail_view div.prd_brand_area div.brnad_link_prd h1.kor`
	imageSelector = overviewSelector +  `div.item_img_view div.detail_img_area ruler-swiper-container div.prd_swiper div.swiper-container div.swiper-wrapper ruler-swiper-slide.swiper-slide div.imgbx div.imgin img`
)

func Scrap(url string) (*item.Item, error) {
	ctx, cancel := repository.NewContext()
	defer cancel()
	
	item := &item.Item{
		Url: url,
	}
	if err := chromedp.Run(
		ctx,
		chromedp.Emulate(device.IPhoneXR),
		chromedp.Navigate(url),
		
		chromedp.WaitVisible(titleSelector),
		
		chromedp.Text(titleSelector, &item.Title),
		chromedp.Text(brandSelector, &item.Brand),
		chromedp.AttributeValue(imageSelector, "data-blazy", &item.ImageSrc, nil),
		
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
				if attrs[i] == "data-blazy" {
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