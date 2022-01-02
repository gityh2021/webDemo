package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"webDemo/models"
)

func main() {
	models.InitDB()
	UserApiRun()
}

func UserApiRun() {
	r := gin.Default()
	r.POST("/user/register", func(c *gin.Context) {
		var user models.User
		// 将request的body中的数据，自动按照json格式解析到结构体
		if err := c.ShouldBindJSON(&user); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, err := user.Register()
		if err != nil{
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"success": user})
	})
	r.GET("/user/:id", func(c *gin.Context){
		// 解析uri上的参数
		idJson := c.Param("id")
		id, err := strconv.Atoi(idJson)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user := models.GetUserById(id)
		c.JSON(http.StatusOK, gin.H{"success": user})
	})
	r.Run()
}