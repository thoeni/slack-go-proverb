package main

import (
	"html"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type proverb struct {
	quote string
	url   string
}

func randomProverb() (proverb, error) {
	// Request the HTML page.
	res, err := http.Get("https://raw.githubusercontent.com/go-proverbs/go-proverbs.github.io/master/index.html")
	if err != nil {
		return proverb{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return proverb{}, err
	}

	var proverbs []proverb

	// Find the review items
	doc.Find("h2 a").Each(func(i int, s *goquery.Selection) {
		quote, _ := s.Html()
		url, _ := s.Attr("href")
		proverbs = append(proverbs, proverb{html.UnescapeString(quote), url})
	})

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	return proverbs[r.Intn(len(proverbs))], nil
}
