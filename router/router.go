package router

import (
    "github.com/gin-gonic/gin"
    "github.com/jiujuanfeng/yunheblog/controller/admin"
    "github.com/jiujuanfeng/yunheblog/utils"
    "html/template"
)

func Route(router *gin.Engine) {
    funcMap := template.FuncMap{
        "dateFormat" : utils.DateFormat,
    }
    router.SetFuncMap(funcMap)
    router.Static("/static","./static")
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
