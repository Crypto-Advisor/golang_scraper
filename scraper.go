package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/groovili/gogtrends"
)

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {
	ctx := context.Background()

	explore, err := gogtrends.Explore(ctx,
		&gogtrends.ExploreRequest{
			ComparisonItems: []*gogtrends.ComparisonItem{
				{
					Keyword: "Go",
					Geo:     "US",
					Time:    "today 12-m",
				},
			},
			Category: 31, // Programming category
			Property: "",
		}, "EN")
	fmt.Println(err)
	fmt.Println(explore)

	c := colly.NewCollector(
		colly.AllowedDomains("amazon.com"),
	)

	c.OnHTML("div.s-result-item", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.Visit("https://www.amazon.com/s?k=phones&crid=3OUB06SCTXYEH&sprefix=phone%2Caps%2C190&ref=nb_sb_noss_1")

	log.Printf("Scraping finished, check for results\n")
}
