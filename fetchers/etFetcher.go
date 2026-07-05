package fetchers

import (
	"slices"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
);

var etLinks []string;
var etTitles []string;
var etDescs []string;

var etNews []News;

func ETFetcher() (news []News) {
	etLinks = []string{};
	etTitles = []string{};
	etDescs = []string{};

	etNews = []News{};

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://bfsi.economictimes.indiatimes.com/"},
		ParseFunc: etFetch,
		RobotsTxtDisabled: true,
	}).Start();
	
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: etLinks,
		ParseFunc: func (g *geziyor.Geziyor, r *client.Response) {
			etScrapeMore(r);
		},
		RobotsTxtDisabled: true,
	}).Start();

	for i, link := range etLinks {
		etNews = append(etNews, News{etTitles[i], etDescs[i], link});
	}

	return etNews;
}

func etFetch(g *geziyor.Geziyor, r *client.Response) {
	parsed := r.HTMLDoc;

	parsed.Find("div.top-story-panel a").Each(func(_ int, a *goquery.Selection) {
		link, _ := a.Attr("href");
		title := strings.TrimSpace(a.Text());
		etLinks = append(etLinks, "https://bfsi.economictimes.indiatimes.com" + link);
		etTitles = append(etTitles, title);
		etDescs = append(etDescs, "");
	});
}

func etScrapeMore(r *client.Response) {
	parsed := r.HTMLDoc;

	link := r.Request.URL.String();

	desc := parsed.Find("span.detail_synopsis").Text();
	etDescs[slices.Index(etLinks, link)] = strings.TrimSpace(desc);
}