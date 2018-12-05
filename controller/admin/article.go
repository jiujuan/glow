package admin

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func ListArticle(c *gin.Context) {
    c.HTML(http.StatusOK, "admin/articlelist.html", gin.H{})
}

func AddArticle(c *gin.Context) {
    c.HTML(http.StatusOK, "admin/articleadd.html", nil)
}
