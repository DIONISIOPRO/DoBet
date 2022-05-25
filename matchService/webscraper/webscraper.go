package webscraper

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

type Webscraper struct{
	BaseUrl string
	chromedp chromedp.Conn
}

func NewScraper(url string) *Webscraper{
	return &Webscraper{
		BaseUrl: url,
	}
}

func (w Webscraper)GetAvailableCountries() []string{
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, time.Second * 15)
	defer cancel()

	var countries []string

	err := chromedp.Run(ctx,
		chromedp.Navigate(w.BaseUrl),
		chromedp.WaitVisible(`body > footer`),
		chromedp.Click(`#example-After`, chromedp.NodeVisible),
		chromedp.Value(`#example-After textarea`, ),
	)
	if err != nil {
		log.Fatal(err)
	}
//	log.Printf("Go's time.After example:\n%s", example)

	return countries
}

func (w Webscraper)GetAvailableLeaguesByCountry(){
	
}

func(w Webscraper) GetAvailableMatchesByLeague(){

}