package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

// Create your own credentials by creating an oauth app at developer.twitter.com
const accessToken = "<<YOUR ACCESS TOKEN>>"
const accessTokenSecret = "<<YOUR ACCESS TOKEN SECRET>>"
const twitterKey = "<<YOUR TWITTER KEY>>"
const twitterSecret = "<<YOUR TWITTER SECRET>>"

func main() {
	var handle string
	if len(os.Args) < 2 {
		fmt.Printf("Error, you must provide at least one handle\n")
		os.Exit(100)
	}
	var handles = os.Args[1:]

	fmt.Printf("Searching tweets by %v\n\n", handles)
	api := anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, twitterKey, twitterSecret)
	for _, handle = range handles {
		printTweet(handle, api)
	}
}

// Prints the latest tweet from a specified twitter handle
func printTweet(handle string, api *anaconda.TwitterApi) {
	v := url.Values{}
	v.Set("count", "1")
	searchResult, _ := api.GetSearch("from:"+handle, v)
	for _, tweet := range searchResult.Statuses {
		fmt.Printf("%v on %v\n%v\n\n", handle, tweet.CreatedAt, tweet.Text)
	}

}
