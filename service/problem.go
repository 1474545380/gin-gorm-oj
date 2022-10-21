package service

import (
	"gin-gorm-oj/define"
	"gin-gorm-oj/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList
// @Summary 问题列表
// @Schemes
// @Param page query int false "请输入当前页，默认第一页"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /getproblemlist [get]
func GetProblemList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	size, err := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Println("get problem list page strconv error:", err)
		return
	}
	//if page===1==>==0
	page = (page - 1) * size
	var count int64
	keyword := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")
	data := make([]*models.Problem, 0)
	tx := models.GetProblemList(keyword, categoryIdentity)

	err = tx.Count(&count).Omit("content").Offset(page).Limit(size).Find(&data).Error
	if err != nil {
		log.Println("Get problem List err:", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  data,
			"count": count,
		},
	})
}
