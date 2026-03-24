package handler

import (
	"go_demo/middleware"
	"go_demo/pkg/response"
	"go_demo/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.Error(c, 400, "参数错误:"+err.Error())
		return
	}

	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	err := service.CreateUser(req.Username, string(hashPwd), req.Email)
	if err != nil {
		response.Error(c, 500, "注册失败："+err.Error())
		return
	}

	response.Success(c, gin.H{"msg": "注册成功"})
}

func UserLogin(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		response.Error(c, 404, "参数错误:"+err.Error())
		return
	}

	user, err := service.GetUserByEmail(req.Email)
	if err != nil {
		response.Error(c, 404, "电邮不存在")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		response.Error(c, 401, "密码错误")
		return
	}

	token, _ := middleware.GenerateToken(user.ID, user.Username)

	response.Success(c, gin.H{"token": token, "username": user.Username})
}

func GetUserInfo(c *gin.Context) {
	userID := c.GetUint("userID")

	user, err := service.GetUserByID(userID)
	if err != nil {
		response.Error(c, 500, "获取用户信息失败")
		return
	}

	response.Success(c, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})

}
