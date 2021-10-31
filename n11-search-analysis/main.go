package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

type Product struct {
	title string
	subTitle string
}

type Products struct {
	products  []Product
}

func main()  {
	doc, err := goquery.NewDocument("https://www.n11.com/arama?q=iphone+13+k%C4%B1l%C4%B1f")
	if err != nil {
		fmt.Println(err)
	}

	var product []Product

	counter := 0

	doc.Find(".column").Each(func(i int, s *goquery.Selection) {


		product[counter].title = s.Find(".productName").Text()
		product[counter].subTitle = s.Find(".proSubTitle").Text()

		counter = counter + 1

	})

	fmt.Println(product)
}
