package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	doc, err := goquery.NewDocument("http://www.borakasmer.com/category/code-first/")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".container .post-row").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".post-title").Text()
		fmt.Println(title)
	})
}
