package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

type Article struct {
	Title string
}

func main() {
	c := colly.NewCollector()

	var article Article

	c.OnHTML("h3.css-11ytn36 a", func(e *colly.HTMLElement) {
		article.Title = e.Text
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.nytimes.com/section/sports")


	fmt.Printf("First article title: %s\n", article.Title)

	jsonData, err := json.Marshal(article)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSON: %s\n", jsonData)

	err = ioutil.WriteFile("output.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
