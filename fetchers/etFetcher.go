package fetchers

import (
    "io"
)

func ETFetcher() (news []News) {
    resp, err := httpClient.Get("https://bfsi.economictimes.indiatimes.com/rss/topstories")
    if err != nil {
        return nil
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil
    }
    news = parse(body)

    return news
}