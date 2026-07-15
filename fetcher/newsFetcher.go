package fetcher

import (
	"encoding/xml"
	"html"
	"io"
	"sync"
	"time"
)

func newsFetch(link string, source string, wg *sync.WaitGroup, ch chan <- []News) {
	defer wg.Done()

	resp, err := httpClient.Get(link)
    if err != nil {
        ch <- []News{}
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        ch <- []News{}
    }

	var feed rssFeed
	if err := xml.Unmarshal(body, &feed); err != nil {
		ch <- []News{}
	}

	news := make([]News, 0, len(feed.Items))

	now := time.Now()

	for _, item := range feed.Items {
		if isTodayOrYesterday(item.PubDate, now) {
			news = append(news, News{
				Title: html.UnescapeString(item.Title),
				Desc:  html.UnescapeString(item.Description),
				Link:  item.Link,
				Source: source,
			})
		}
	}

	ch <- news
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