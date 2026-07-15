package fetcher

import (
	"net/http"
	"time"
)

type News struct {
	Title	string	`json:"title"`
	Desc	string	`json:"desc"`
	Link	string	`json:"link"`
	Source	string	`json:"source"`
}

type rssFeed struct {
	Items		[]rssItem `xml:"channel>item"`
}


type rssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate		string `xml:"pubDate"`
}

var httpClient = &http.Client{
	Timeout:	10 * time.Second,
	Transport:	&http.Transport{
		MaxIdleConns:			100,
		MaxIdleConnsPerHost:	10,
		MaxConnsPerHost:		10,
		IdleConnTimeout:		90 * time.Second,
	},
}