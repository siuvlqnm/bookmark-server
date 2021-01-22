package main_test

import (
	"testing"

	"github.com/siuvlqnm/bookmark/service"
	"github.com/siuvlqnm/bookmark/utils"
)

func TestGetWebInfo(T *testing.T) {
	url := "http://test.com"
	err, p := utils.ParseUrl(url)
	if err != nil {
		T.Fatal(err)
	}

	err, _ = service.GetWebSite(p.Domain)
	// if err != nil {
	// 	T.Log(w.ID)
	// }
}
