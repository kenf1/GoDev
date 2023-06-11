package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

// create struct w/ specified json fields
type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"),
	)

	var items []item

	c.OnHTML("div.col-sm-9 div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		item := item{
			Name:   h.ChildText("h2.product-title"),
			Price:  h.ChildText("div.sale-price"),
			ImgUrl: h.ChildAttr("img", "src"),
		}

		//print all product names on page
		// fmt.Println(h.ChildText("h2.product-title"))

		// fmt.Println(item)

		items = append(items, item)
	})

	c.Visit("https://j2store.net/v3/index.php/shop/clothing/t-shirts")

	//print items
	fmt.Println(items)

	//save to json
	file, err := json.Marshal(items)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("items.json", file, 0644)
}
