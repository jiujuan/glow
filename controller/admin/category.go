package admin

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "time"
    "github.com/jiujuan/yunheblog/model"
    "github.com/jiujuan/yunheblog/db"
    log "github.com/sirupsen/logrus"
)

const (
    CategoryAddTpl = "admin/categoryadd.html"
    CategoryIndexTpl = "admin/category.html"
)

func IndexCategory(c *gin.Context) {
    var catList []model.Category
    err := db.DB.Order("id desc").Find(&catList).Error
    if err != nil {
        log.Errorf("get category error: %s ", err.Error())
    }
    c.HTML(http.StatusOK, CategoryIndexTpl, gin.H{
        "catList": catList,
    })
}

func AddCategory(c *gin.Context) {
    c.HTML(http.StatusOK, CategoryAddTpl, nil)
}

func EditCategory(c *gin.Context) {
    id := c.Param("id")
    var catInfo model.Category
    err := db.DB.First(&catInfo, id).Error
    if err != nil {
        log.Errorf("get category error: %s, id: %d ", err.Error(), id)
    }
    c.HTML(http.StatusOK, CategoryAddTpl, gin.H{
        "catInfo": catInfo,
    })
}

func SaveCategory(c *gin.Context) {
    catName := c.PostForm("cat_name")
    catInfo := c.PostForm("cat_info")
    catKey := c.PostForm("cat_key")

    ID, _ := strconv.Atoi(c.PostForm("id"))
    updateTime := 0
    if ID > 0 {
        updateTime = int(time.Now().Unix())
    }

    categoryModel := model.Category{
        CatName: catName,
        CatInfo: catInfo,
        CatKey: catKey,
        BaseModel: model.BaseModel{
            CreateTime: int(time.Now().Unix()),
            UpdateTime: updateTime,
        },
    }

    var err error
    if ID > 0 {
        err = db.DB.Model(&categoryModel).Where("id = ?", ID).Updates(categoryModel).Error
    } else {
        err = db.DB.Create(&categoryModel).Error
    }
    if err == nil {
        c.Redirect(http.StatusMovedPermanently, "/admin/category")
    } else {
        c.HTML(http.StatusOK, CategoryAddTpl, nil)
    }
}

func DeleteCategory(c *gin.Context) {
    id,_ := strconv.Atoi(c.Param("id"))

    catModel := model.Category{}
    catModel.ID = uint(id)

    var err error
    log.Info(id)
    if uint(id) > 0 {
        err = db.DB.Delete(&catModel).Error
    }

    c.JSON(http.StatusOK, gin.H{
        "msg": err==nil,
        "code": 200,
    })
}

func Error() string {
    return "cat error"
}
