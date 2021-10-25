package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
)

type Credentials struct {
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessTokenSecret string
}

func main()  {
	creds := Credentials{
		AccessToken:"access-token",
		AccessTokenSecret:"access-token-secret",
		ConsumerKey:"consumer-key",
		ConsumerSecret:"consumer-secret",
	}
	
	client, err := getClient(&creds)

	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}
	
	tweet, resp, err := client.Statuses.Update("Golang app test tweet #golang", nil)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode==200 {
		fmt.Println(tweet.ID)
	}
}

func getClient(creds *Credentials) (*twitter.Client, error) {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)

	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	return client, nil
}