package service

import (
	"gin-gorm-oj/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProblemList(c *gin.Context) {
	models.GetProblemList()
	c.String(http.StatusOK, "Get Problem List")
}
