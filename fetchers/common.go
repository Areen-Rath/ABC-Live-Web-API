package fetchers

import (
	"encoding/xml"
	"html"
	"net/http"
	"time"
)

type News struct {
	Title	string	`json:"title"`
	Desc	string	`json:"desc"`
	Link	string	`json:"link"`
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


func parse(body []byte) (news []News) {
	var feed rssFeed
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil
	}

	news = make([]News, 0, len(feed.Items))

	now := time.Now()

	for _, item := range feed.Items {
		if isTodayOrYesterday(item.PubDate, now) {
			news = append(news, News{
				Title: html.UnescapeString(item.Title),
				Desc:  html.UnescapeString(item.Description),
				Link:  item.Link,
			})
		}
	}

	return news
}

func isTodayOrYesterday(pubDate string, now time.Time) bool {
	datetime, err := time.Parse(time.RFC1123Z, pubDate)
	if err != nil {
		return false
	}

	date := datetime.Truncate(24 * time.Hour)
	today := now.Truncate(24 * time.Hour)
	yesterday := now.AddDate(0, 0, -1).Truncate(24 * time.Hour)
	return time.Time.Equal(date, today) || time.Time.Equal(date, yesterday)
}