package main

import (
	"github.com/gin-gonic/gin"
	"github.com/project_login/controlls"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/signup", controlls.SignupPage)
	r.POST("/signup", controlls.SignupUser)
	r.GET("/login", controlls.Loginpage)
	r.POST("/login", controlls.Loginuser)
	r.GET("/home", controlls.Homepage)
	r.GET("/adminloginpage", controlls.Adminloginpage)
	r.POST("/adminloginpage", controlls.Adminlogin)
	r.GET("/adminpanel", controlls.Adminpanel)
	r.POST("/adminpanel", controlls.Search, controlls.Delete)

	r.Run(":6060")

}
