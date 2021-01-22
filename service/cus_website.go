package service

import (
	"errors"

	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
	"gorm.io/gorm"
)

func CreateWebSite(w *model.CusWebsite) (err error, webInfo *model.CusWebsite) {
	var website model.CusWebsite
	if !errors.Is(global.GVA_DB.Where("domain = ?", w.Domain).First(&website).Error, gorm.ErrRecordNotFound) {
		return err, &website
	}
	err = global.GVA_DB.Create(&w).Error
	return err, w
}

func GetWebSite(domain string) (err error, w *model.CusWebsite) {
	var website model.CusWebsite
	if errors.Is(global.GVA_DB.Where("domain = ?", domain).First(&website).Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound, &website
	}
	return err, &website
}
