package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("h1", func(element *colly.HTMLElement) {
		fmt.Println("\n-------------------------")
		fmt.Println("Car: ", element.Text)
		fmt.Println("-------------------------")
	})
	c.OnHTML(".amount--3NTpl", func(element *colly.HTMLElement) {
		fmt.Println("Price: ", element.Text)
		fmt.Println("-------------------------")
	})
	c.OnHTML(".sub-title--37mkY", func(element *colly.HTMLElement) {
		fmt.Println("Posted On: ", element.Text)
		fmt.Println("-------------------------")
	})
	c.OnHTML(".two-columns--19Hyo", func(element *colly.HTMLElement) {
		fmt.Println(element.Text)
	})
	desc := false
	c.OnHTML(".description--1nRbz > p", func(element *colly.HTMLElement) {
		if !desc {
			fmt.Println("-------------------------")
			fmt.Println("Description: ", element.Text)
		} else {
			fmt.Println(element.Text)
		}
		desc = true
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("\nFinished", r.Request.URL)
	})
	_ = c.Visit("https://ikman.lk/en/ad/nissan-sunny-fb15-2001-for-sale-kandy-8")
}
