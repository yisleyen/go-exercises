package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type currency struct {
	XMLName xml.Name `xml:"STOCK"`
	DESC    string   `xml:"DESC"`
	LAST    string   `xml:"LAST"`
	PERNC   string   `xml:"PERNC"`
}

type market struct {
	XMLName    xml.Name   `xml:"ICPIYASA"`
	Currencies []currency `xml:"STOCK"`
}

func (c currency) String() string {
	return fmt.Sprintf("Sembol: %s - Son Değer: %s - Değişim: %s \n", c.DESC, c.LAST, c.PERNC)
}

func main() {
	resp, err := http.Get("http://realtime.paragaranti.com/asp/xml/icpiyasa.asp")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: ", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	var result market
	xml.Unmarshal(data, &result)

	fmt.Println(result.Currencies)
}
