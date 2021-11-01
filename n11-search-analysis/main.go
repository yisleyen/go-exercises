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

	var product Product

	doc.Find(".column").Each(func(i int, s *goquery.Selection) {

		product.title = s.Find(".productName").Text()
		product.subTitle = s.Find(".proSubTitle").Text()


	})

	fmt.Println(product)
}
