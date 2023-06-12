package main

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.ittf.com"),
	)

	raw_data := []string{}

	//extract + print html table contents
	c.OnHTML("tbody tr.rrow.age-SEN", func(e *colly.HTMLElement) {
		e.ForEach("td", func(_ int, h *colly.HTMLElement) {
			// fmt.Println(h.Text)
			raw_data = append(raw_data, h.Text)
		})
	})

	c.Visit("https://www.ittf.com/wp-content/uploads/2023/06/2023_23_SEN_MS.html")

	// fmt.Println(raw_data)
	fmt.Println(reflect.TypeOf(raw_data))

	//convert string to byte
	var bytes []byte
	for _, str := range raw_data {
		bytes = append(bytes, []byte(str)...)
	}

	//save to txt
	err := ioutil.WriteFile("player_info.txt", []byte(bytes), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	//save to json
	// file, err := json.Marshal(player_infos)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// os.WriteFile("player_info.json", file, 0644)
}
