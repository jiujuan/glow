package router

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jiujuan/glow/controller/admin"
	"github.com/jiujuan/glow/utils"
)

// NoResponse 请求的URL不存在，返回404
func NoResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error":  "404, page not exists!",
		"status": 404,
	})
}

//Route 请求的url路由
func Route(router *gin.Engine) {
	funcMap := template.FuncMap{
		"dateFormat": utils.DateFormat,
	}
	router.SetFuncMap(funcMap)
	router.Static("/static", "./static")
	router.LoadHTMLGlob("views/**/*")
	adminRouter := router.Group("/admin")
	{
		adminRouter.GET("/", admin.Index)
		adminRouter.GET("/index", admin.Index)

		adminRouter.GET("/article/list", admin.ListArticle)
		adminRouter.GET("/article/add", admin.AddArticle)

		adminRouter.GET("/sysconfig", admin.SysConfig)
		adminRouter.POST("/sysconfig/save", admin.SaveSysConfig)

		adminRouter.GET("/category", admin.IndexCategory)
		adminRouter.GET("/category/index", admin.IndexCategory)
		adminRouter.GET("/category/add", admin.AddCategory)
		adminRouter.GET("/category/edit/:id", admin.EditCategory)
		adminRouter.POST("/category/save", admin.SaveCategory)
		adminRouter.POST("/category/delete/:id", admin.DeleteCategory)
	}
}
