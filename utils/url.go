package utils

import (
	"net/url"
)

type UrlInfo struct {
	TargetUrl string
	Domain    string
	Path      string
	Query     string
}

func ParseUrl(link string) (err error, I *UrlInfo) {
	u, err := url.Parse(link)
	if err != nil {
		return err, I
	}
	targetUrl := link
	I = &UrlInfo{TargetUrl: targetUrl, Domain: u.Host, Path: u.Path, Query: u.RawQuery}
	return err, I
}
