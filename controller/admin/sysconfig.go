package admin

import (
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/jiujuan/glow/db"
	"github.com/jiujuan/glow/model"
	"net/http"
	"strconv"
	"time"
)

const (
	SysConfigTPL = "admin/sysconfig.html"
)

func SysConfig(c *gin.Context) {
	var sysConfig model.SysConfig
	err := db.DB.First(&sysConfig).Error
	if err != nil {
		log.Errorf("get config error: %s ", err.Error())
	}
	c.HTML(http.StatusOK, SysConfigTPL, gin.H{
		"configInfo": sysConfig,
	})
}

func SaveSysConfig(c *gin.Context) {
	siteName := c.PostForm("site_name")
	keyWord := c.PostForm("key_word")
	siteDesc := c.PostForm("site_desc")
	isOpen := c.PostForm("is_open")
	adminEmail := c.PostForm("admin_email")
	closeDesc := c.PostForm("close_desc")
	isOpenInt, _ := strconv.Atoi(isOpen)
	ID, _ := strconv.Atoi(c.PostForm("id"))

	updateTime := 0
	if ID > 0 {
		updateTime = int(time.Now().Unix())
	}
	sysConfig := model.SysConfig{
		SiteName:   siteName,
		KeyWord:    keyWord,
		SiteDesc:   siteDesc,
		IsOpen:     isOpenInt,
		AdminEmail: adminEmail,
		CloseDesc:  closeDesc,
		BaseModel: model.BaseModel{
			CreateTime: int(time.Now().Unix()),
			UpdateTime: updateTime,
		},
	}

	var err error
	if ID > 0 {
		err = db.DB.Model(&sysConfig).Where("id = ?", ID).Updates(sysConfig).Error
	} else {
		err = db.DB.Create(&sysConfig).Error
	}
	if err == nil {
		c.Redirect(http.StatusMovedPermanently, "/admin/sysconfig")
	} else {
		log.Errorf("insert data error: %s", err.Error())
		c.HTML(http.StatusOK, SysConfigTPL, nil)
	}
}
