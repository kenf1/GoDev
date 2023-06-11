package main

import (
	//core
	"encoding/json"
	"fmt"
	"os"

	//additional
	"github.com/gocolly/colly"
)

// struct (similar to python classes)
type Quote struct {
	Quote  string
	Author string
}

func main() {
	//slice of struct
	quotes := []Quote{}

	//setup colly
	c := colly.NewCollector()
	colly.AllowedDomains("quotes.toscrape.com")

	//print which url colly is visiting
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	//Response code 200 == good to go
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response code =", r.StatusCode)
	})

	//print error if error detected
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error", err.Error())
	})

	//print indiv aspects of quote (all quotes followed by all authors)
	// c.OnHTML(".text", func(h *colly.HTMLElement) {
	// 	fmt.Println("Quote", h.Text)
	// })
	// c.OnHTML(".author", func(h *colly.HTMLElement) {
	// 	fmt.Println("Author", h.Text)
	// })

	//quote follwed by author
	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		div := h.DOM
		quote := div.Find(".text").Text()
		author := div.Find(".author").Text()

		//print quote ()
		// fmt.Println("Quote:", quote, "\nBy:", author, "\n")

		//append to struct
		q := Quote{
			Quote:  quote,
			Author: author,
		}
		quotes = append(quotes, q)
	})

	//add `/random` for single quote, rm for all quotes
	c.Visit("http://quotes.toscrape.com")

	//save to json in working dir (Scrape)
	file, err := os.Create("quotes.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(quotes)
	if err != nil {
		fmt.Println(err)
		return
	}

	//print all quotes
	fmt.Println(quotes)
}
