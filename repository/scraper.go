package repository

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/jmshin92/scraper/closers"
	"github.com/jmshin92/scraper/item"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

type Scraper interface {
	WaitAction() chromedp.Action
	ScrapActions(*item.Item) []chromedp.Action
}

var (
	handler *Handler
)

func init() {
	handler = &Handler{}
	handler.ctx, handler.cancel = NewContext()
	
	closers.AddClosers(handler)
}

func Scrap(url string, scraper Scraper) (*item.Item, error) {
	return handler.scrap(url, scraper)
}

type Handler struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func (this *Handler) scrap(url string, scraper Scraper) (*item.Item, error) {
	item := &item.Item{}
	
	actions := []chromedp.Action {
		//chromedp.Emulate(device.),
		chromedp.Navigate(url),
		scraper.WaitAction(),
	}
	
	actions = append(actions, scraper.ScrapActions(item)...)
	
	if err := chromedp.Run(
		handler.ctx,
		actions...,
	); err != nil {
		return nil, err
	}
	
	return item, nil
}

func (this *Handler) Close() {
	logrus.Info("close scraper handler")
	this.cancel()
}

func NewContext() (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Headless,
		chromedp.DisableGPU,
		chromedp.NoSandbox,

		chromedp.Flag("incognito", true),
		chromedp.Flag("ignore-certificate-errors", true),

		chromedp.Flag("disable-images", true),
		chromedp.Flag("blink-settings", "imagesEnabled=false"),

		chromedp.Flag("disable-background-networking", true),
		chromedp.Flag("enable-features", "NetworkService,NetworkServiceInProcess"),
		chromedp.Flag("disable-background-timer-throttling", true),
		chromedp.Flag("disable-backgrounding-occluded-windows", true),
		chromedp.Flag("disable-breakpad", true),
		chromedp.Flag("disable-client-side-phishing-detection", true),
		chromedp.Flag("disable-default-apps", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-features", "site-per-process,TranslateUI,BlinkGenPropertyTrees"),
		chromedp.Flag("disable-hang-monitor", true),
		chromedp.Flag("disable-ipc-flooding-protection", true),
		chromedp.Flag("disable-popup-blocking", true),
		chromedp.Flag("disable-prompt-on-repost", true),
		chromedp.Flag("disable-renderer-backgrounding", true),
		chromedp.Flag("disable-sync", true),
		chromedp.Flag("force-color-profile", "srgb"),
		chromedp.Flag("metrics-recording-only", true),
		chromedp.Flag("safebrowsing-disable-auto-update", true),
		chromedp.Flag("enable-automation", true),
		chromedp.Flag("password-store", "basic"),
		chromedp.Flag("use-mock-keychain", true),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	ctx, cancel = context.WithTimeout(ctx, 20 * time.Second)
	return ctx, cancel
}
