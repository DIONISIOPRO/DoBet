package data

import (
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func GetDocument(url string) (*goquery.Document, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
