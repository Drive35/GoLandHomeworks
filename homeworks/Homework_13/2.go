package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

var (
	url     = "https://www.zakon.kz/"
)

func main () {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)
}