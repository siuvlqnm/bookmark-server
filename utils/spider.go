package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type WebInfo struct {
	Title       string
	Description string
}

func GetWebInfo(url string) (err error, w *WebInfo) {
	request, _ := http.NewRequest("GET", url, nil)

	request.Header.Set("User-Agent", RandomUserAgent())
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(request)
	if err != nil {
		return err, w
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status)), w
	}

	// // Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err, w
	}

	t := doc.Find("title").Text()
	d := doc.Find("meta").Map(func(i int, s *goquery.Selection) string {
		var d string
		if name, _ := s.Attr("name"); strings.EqualFold(name, "description") {
			d, _ := s.Attr("content")
			return d
		}
		return d
	})
	for index, v := range d {
		if d[index] != "" {
			return err, &WebInfo{Title: t, Description: v}
		}
	}
	return err, &WebInfo{Title: t}
}

var uaGens = []func() string{
	genFirefoxUA,
	genChromeUA,
}

func RandomUserAgent() string {
	return uaGens[rand.Intn(len(uaGens))]()
}

var ffVersions = []float32{
	58.0,
	57.0,
	56.0,
	52.0,
	48.0,
	40.0,
	35.0,
}

var chromeVersions = []string{
	"65.0.3325.146",
	"64.0.3282.0",
	"41.0.2228.0",
	"40.0.2214.93",
	"37.0.2062.124",
}

var osStrings = []string{
	"Macintosh; Intel Mac OS X 10_10",
	"Windows NT 10.0",
	"Windows NT 5.1",
	"Windows NT 6.1; WOW64",
	"Windows NT 6.1; Win64; x64",
	"X11; Linux x86_64",
}

func genFirefoxUA() string {
	version := ffVersions[rand.Intn(len(ffVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s; rv:%.1f) Gecko/20100101 Firefox/%.1f", os, version, version)
}

func genChromeUA() string {
	version := chromeVersions[rand.Intn(len(chromeVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", os, version)
}
