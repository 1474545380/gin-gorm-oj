package service

import (
	"fmt"
	"gin-gorm-oj/help"
	"gin-gorm-oj/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// GetUserDetail
// @Tags 用户方法
// @Summary 用户详情
// @Param identity query string false "user identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /userdetail [get]
func GetUserDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户唯一标识不能为空",
		})
		return
	}
	data := new(models.User)
	err := models.DB.Omit("password").Where("identity = ?", identity).Find(&data).Error
	if data.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "没有查询到该用户",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "get user detail by identity:" + identity + " Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

// Login
// @Tags 用户方法
// @Summary 用户登陆
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /login [post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if password == "" || username == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "必填信息为空",
		})
		return
	}
	password = help.GetMd5(password)
	fmt.Println("password:" + password)
	data := new(models.User)
	err := models.DB.Where("name = ? AND password = ?", username, password).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "用户名或密码错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "get user err:" + err.Error(),
		})
		return
	}
	token, err := help.GenerateToken(data.Identity, data.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "token生成错误" + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

// SendCode
// @Tags 用户方法
// @Summary 发送验证码
// @Param email formData string true "email"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /sendcode [post]
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数不正确",
		})
		return
	}
	code := "123123"
	err := help.SendCode(email, code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "send code error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "发送成功",
	})
}
