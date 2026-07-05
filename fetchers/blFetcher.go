package fetchers;

import (
	"fmt"
	"slices"
	"strings"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/PuerkitoBio/goquery"
);

var blLinks []string;
var blTitles []string;
var blDescs []string;

var blNews []News;

func BLFetcher() (news []News) {
	blLinks = []string{};
	blTitles = []string{};
	blDescs = []string{};

	blNews = []News{};

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://www.thehindubusinessline.com/money-and-banking"},
		ParseFunc: blFetch,
		RobotsTxtDisabled: true,
	}).Start();

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: blLinks,
		ParseFunc: func (g *geziyor.Geziyor, r *client.Response) {
			blScrapeMore(r);
		},
		RobotsTxtDisabled: true,
	}).Start();

	for i, link := range blLinks {
		if blDescs[i] != "" {
			blNews = append(blNews, News{blTitles[i], blDescs[i], link});
		}
	}

	fmt.Println(blNews);
	return blNews;
}

func blFetch(g *geziyor.Geziyor, r *client.Response) {
	parsed := r.HTMLDoc;

	parsed.Find("div.right-content a").Each(func (_ int, a *goquery.Selection) {
		link, _ := a.Attr("href");
		title := a.Text();
		if(!slices.Contains(blLinks, link) &&
			link[36:47] != "/companies/") {
			blLinks = append(blLinks, link);
			blTitles = append(blTitles, title);
			blDescs = append(blDescs, "");
		}
	});
}

func blScrapeMore(r *client.Response) {
	parsed := r.HTMLDoc;

	link := r.Request.URL.String();

	desc := parsed.Find("h2.sub-title").First().Text();
	if !strings.Contains(desc, "LIVE") {
		blDescs[slices.Index(blLinks, link)] = strings.TrimSpace(desc);
	}
}