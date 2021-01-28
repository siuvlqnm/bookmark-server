package cus

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/model"
	"github.com/siuvlqnm/bookmark/model/request"
	"github.com/siuvlqnm/bookmark/model/response"
	"github.com/siuvlqnm/bookmark/service"
	"github.com/siuvlqnm/bookmark/utils"
)

func GetWebInfo(c *gin.Context) {
	var u request.Website
	_ = c.ShouldBindJSON(&u)

	if err := utils.Verify(u, utils.GetWebInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err, p := utils.ParseUrl(u.Url)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err, gws := service.GetWebSite(p.Domain)
	if err == nil {
		response.OkWithData(response.WebsiteResponse{TargetUrl: p.TargetUrl, Protocol: gws.Protocol, Domain: gws.Domain, Path: p.Path, Query: p.Query, Title: gws.Title, Description: gws.Description}, c)
		return
	}

	if err, gwi := utils.GetWebInfo(u.Url); err != nil {
		website := &model.CusWebsite{Protocol: p.Protocol, Domain: p.Domain}
		service.CreateWebSite(website)
		response.FailWithDetailed(response.WebsiteResponse{TargetUrl: p.TargetUrl, Protocol: p.Protocol, Domain: p.Domain, Path: p.Path, Query: p.Query}, err.Error(), c)
		return
	} else {
		website := &model.CusWebsite{Protocol: p.Protocol, Domain: p.Domain, Title: gwi.Title, Description: gwi.Description}
		service.CreateWebSite(website)
		response.OkWithData(response.WebsiteResponse{TargetUrl: p.TargetUrl, Protocol: p.Protocol, Domain: p.Domain, Path: p.Path, Query: p.Query, Title: gwi.Title, Description: gwi.Description}, c)
		return
	}
}

// func GetWebInfos(c *gin.Context) {
// 	var u request.Website
// 	_ = c.ShouldBindJSON(&u)

// 	if err := utils.Verify(u, utils.GetWebInfoVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}

// 	err, p := utils.ParseUrl(u.Url)
// 	if err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}

// 	websiteInfo := make(map[string]interface{})
// 	_, err = utils.GetSetValue("website", p.Domain)
// 	if err == nil {
// 		val, _ := utils.GetHashValue("website", p.Domain)
// 		response.OkWithData(response.WebsiteResponse{TargetUrl: p.TargetUrl, Protocol: val["protocol"], Domain: val["domain"], Path: p.Path, Query: p.Query, Title: val["title"], Description: val["description"]}, c)
// 		return
// 	}

// 	_, gwi := utils.GetWebInfo(u.Url)
// 	websiteInfo["protocol"] = p.Protocol
// 	websiteInfo["domain"] = p.Domain
// 	websiteInfo["title"] = gwi.Title
// 	websiteInfo["description"] = gwi.Description
// 	utils.SetHashValue("website", p.Domain, websiteInfo)

// 	website := &model.CusWebsite{Protocol: p.Protocol, Domain: p.Domain, Title: gwi.Title, Description: gwi.Description}
// 	err, _ = service.GetWebSite(p.Domain)
// 	if err != nil {
// 		service.CreateWebSite(website)
// 	} else {
// 		service.UpdateWebSite(website)
// 	}
// 	response.OkWithData(response.WebsiteResponse{TargetUrl: p.TargetUrl, Protocol: p.Protocol, Domain: p.Domain, Path: p.Path, Query: p.Query, Title: gwi.Title, Description: gwi.Description}, c)
// 	return
// }
