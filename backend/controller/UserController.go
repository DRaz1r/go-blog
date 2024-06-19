/*
*
@author: Azir
@desc:
@date: 6/19/24
*
*/
package controller

import (
	"backend/common"
	"backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	db := common.GetDB()
	var requestUser model.User
	c.Bind(&requestUser)
	userName := requestUser.UserName
	phoneNumber := requestUser.PhoneNumber
	password := requestUser.Password
	// 验证
	var user model.User
	db.Where("phone_number = ?", phoneNumber).First(&user)
	if user.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "用户已存在",
		})
		return
	}
	// 密码加密
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 创建用户
	newUser := model.User{
		UserName:    userName,
		PhoneNumber: phoneNumber,
		Password:    string(hashPassword),
		Avatar:      "/images/default_avatar.jpg",
		Collects:    model.Array{},
		Following:   model.Array{},
		Fans:        0,
	}
	db.Create(&newUser)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func Login(c *gin.Context) {
	db := common.GetDB()
	var requestUser model.User
	c.Bind(&requestUser)
	phoneNumber := requestUser.PhoneNumber
	password := requestUser.Password
	// 验证
	var user model.User
	db.Where("phone_number = ?", phoneNumber).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "用户不存在",
		})
		return
	}
	// 密码验证
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "密码错误",
		})
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": token,
		},
	})
	return
}

// GetInfo 登录后获取信息
func GetInfo(c *gin.Context) {
	// 获取上下文中的用户信息
	user, _ := c.Get("user")
	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"id": user.(model.User).ID, "avatar": user.(model.User).Avatar},
		"msg":  "登录获取信息成功",
	})
}
