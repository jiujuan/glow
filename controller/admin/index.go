package admin

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Index(c *gin.Context) {
    c.HTML(http.StatusOK, "admin/index.html", nil)
}
