package main

import (
	"fmt"
	"net/http"
	"service/user"
	"utils"

	"github.com/gin-gonic/gin"
)

//next work:
//login
//auth
//redis
//distribute

func main() {
	log := utils.GetLog("main")
	log.Println("init service...")
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.Static("/res", "../public")
	g.LoadHTMLGlob("../tmpl/**/*")

	//index page
	g.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "gets/index", gin.H{})
	})

	//signup page
	g.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "gets/signup", gin.H{})
	})

	//do signup
	g.POST("/signup", func(c *gin.Context) {
		var userJson user.UserEntity
		err := c.Bind(&userJson)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("input userJson,%v", userJson)
		//res := userJson.QueryUser()
		count, err := userJson.InsertUser(&userJson)
		if err != nil || count < 1 {
			c.JSON(http.StatusOK, gin.H{"msg": "failed"})
		} else {
			c.JSON(http.StatusOK, gin.H{"msg": "ok"})
		}
	})

	g.Run(":3000")
}
