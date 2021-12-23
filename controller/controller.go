package controller

import (
	"fmt"
	"kaoqin/common"
	"math/rand"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"kaoqin/user"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	name := ctx.PostForm("name")
	pwd := ctx.PostForm("password")
	if len(pwd) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  fmt.Sprintf("密码不得小于6位,当前用户名:%s ,密码:%s", name, pwd),
		})
		return
	}
	if len(name) == 0 {
		name = randomString(10)
	}
	var hasUser user.User
	res := db.First(&hasUser, "name = ?", name)
	if res.RowsAffected > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "已有用户名",
		})
		return
	}

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": "422",
			"msg":  "哈希密码错误",
		})
	}
	newUser := user.User{
		Name:     name,
		Password: string(hashPwd),
	}

	db.Create(&newUser)
	ctx.JSON(http.StatusOK, gin.H{
		"name": name,
		"msg":  "注册成功",
	})
}

func Login(ctx *gin.Context) {
	db := common.GetDB()
	pwd := ctx.PostForm("password")
	name := ctx.PostForm("name")
	if len(pwd) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "密码不能小于6位",
		})
		return
	}
	var res user.User
	db.First(&res, "name = ?", name)
	if bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(pwd)) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}
	tokenString, err := common.ReleaseJWT(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "token创建失败",
		})
		return
	}
	fmt.Println("this is token: ", tokenString)
	tokenString = "Bearer " + tokenString
	ctx.Header("Authorization", tokenString)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登陆成功",
	})
}

func Info(ctx *gin.Context) {
	u, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": 401, "msg": "用户转换失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"name": u.(user.User).Name,
	})
}

func randomString(n int) string {
	letters := []byte("qwertyuiopasdfghjklzxcvbnm")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
