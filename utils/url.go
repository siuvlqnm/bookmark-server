package utils

import (
	"net/url"
)

type UrlInfo struct {
	TargetUrl string
	Protocol  string
	Domain    string
	Path      string
	Query     string
}

func ParseUrl(link string) (err error, I *UrlInfo) {
	u, err := url.Parse(link)
	if err != nil {
		return err, I
	}
	targetUrl := u.Scheme + "://" + u.Host + u.Path + "?" + u.RawQuery
	I = &UrlInfo{TargetUrl: targetUrl, Protocol: u.Scheme, Domain: u.Host, Path: u.Path, Query: u.RawQuery}
	return err, I
}
