package fetchers

import (
    "io"
)

func BLFetcher() (news []News) {
    resp, err := httpClient.Get("https://www.thehindubusinessline.com/money-and-banking/feeder/default.rss")
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