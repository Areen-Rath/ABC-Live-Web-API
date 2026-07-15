package fetcher

import "sync"

func Fetch() (et []News, bl []News, rbi [2][7]string) {
    var wg sync.WaitGroup

    etNews := make(chan []News)
    blNews := make(chan []News)
    stats := make(chan [2][7]string)

    wg.Add(3)

    go newsFetch("https://bfsi.economictimes.indiatimes.com/rss/topstories", "Economic Times BFSI", &wg, etNews)
    go newsFetch("https://www.thehindubusinessline.com/money-and-banking/feeder/default.rss", "Business Line", &wg, blNews)
    go RBIFetch(&wg, stats)

    go func() {
        wg.Wait()
        close(etNews)
        close(blNews)
        close(stats)
    }()

    et = <- etNews
    bl = <- blNews
    rbi = <- stats
    return et, bl, rbi
}