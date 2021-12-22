package main

import (
	"fmt"
	"kaoqin/User"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()

	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
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
		newUser := User.User{
			Name:     name,
			Password: pwd,
		}
		db.Create(newUser)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	})
	r.Run()
}

func initDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/kaoqin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect database,err :" + err.Error())
	}
	db.AutoMigrate(&User.User{})
	return db
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
