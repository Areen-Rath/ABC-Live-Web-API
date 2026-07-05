package fetchers;

import (
	"strings"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/PuerkitoBio/goquery"
);

var data [7]string;
var fyph, rr, br, rrr, crr, slr, usd, fyp string;

var headers = [7]string{
	"Policy Repo Rate",
	"Bank Rate",
	"Fixed Reverse Repo Rate",
	"CRR",
	"SLR",
	"INR / 1 USD",
	"",
};

var mix [2][7]string;

func RBIFetcher() (mix [2][7]string) {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://www.rbi.org.in/"},
		ParseFunc: rbiFetch,
		RobotsTxtDisabled: true,
	}).Start();

	mix[0] = headers;
	mix[1] = data;

	return mix;
}

func rbiFetch(g *geziyor.Geziyor, r *client.Response) {
	parsed := r.HTMLDoc;
	
	count := 0;
	parsed.Find("div.accordionContent tr").Each(func (_ int, tr *goquery.Selection) {
		value := strings.TrimSpace(strings.TrimSpace(tr.Find("td").Text())[1:]);
		switch count {
			case 0:
				rr = value;
				data[0] = rr;
			case 3:
				br = value;
				data[1] = br;
			case 4:
				rrr = value;
				data[2] = rrr;
			case 5:
				crr = value;
				data[3] = crr;
			case 6:
				slr = value;
				data[4] = slr;
			case 7:
				usd = value;
				data[5] = usd;
			case 19:
				fyph = strings.TrimSpace(tr.Find("th").Text());
				fyp = value[:len(value) - 2];
				headers[6] = fyph
				data[6] = fyp;
		}
		count++;
	});
}