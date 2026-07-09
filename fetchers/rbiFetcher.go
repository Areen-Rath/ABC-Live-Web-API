package fetchers

import (
    "strings"
	
    "github.com/PuerkitoBio/goquery"
)

func RBIFetcher() (mix [2][7]string) {
	headers := [7]string{}
	data := [7]string{}

	resp, err := httpClient.Get("https://rbi.org.in")
	if err != nil {
		mix[0] = headers
        mix[1] = data
        return mix
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
        mix[0] = headers
        mix[1] = data
        return mix
    }

	count := 0
	foundCount := 0
	positions := map[int]int{0: 0, 3: 1, 4: 2, 5: 3, 6: 4, 7: 5, 19: 6}
	doc.Find("div.accordionContent tr").Each(func(_ int, tr *goquery.Selection) {
		if foundCount == 7 {
			return
		}
		if index, ok := positions[count]; ok && index <= 6 {
			header := strings.TrimSpace(tr.Find("th").Text())
			value := strings.TrimSpace(strings.TrimSpace(tr.Find("td").Text())[1:])
			headers[index] = header
			if index != 6 {
				data[index] = value
			} else {
				data[index] = value[:len(value) - 2]
			}
			foundCount++
		}
		count++
	})

	mix[0] = headers
	mix[1] = data
	return mix
}