package controlls

import (
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/project_login/config"
	"github.com/project_login/models"
)

func SignupPage(c *gin.Context) {
	//Integrating the html Page
	tmpl := template.Must(template.ParseFiles("views/signup.html"))
	tmpl.Execute(c.Writer, nil)

}

func SignupUser(c *gin.Context) {
	c.Request.ParseForm()
	type Userdata struct {
		fname    string
		lname    string
		uname    string
		mail     string
		password string
	}
	var data Userdata
	data.fname = c.Request.PostForm["firstname"][0]
	data.lname = c.Request.PostForm["lastname"][0]
	data.uname = c.Request.PostForm["username"][0]
	data.mail = c.Request.PostForm["email"][0]
	data.password = c.Request.PostForm["password"][0]
	if data.fname == "" || data.lname == "" || data.uname == "" || data.mail == "" || data.password == "" {
		c.Redirect(http.StatusMovedPermanently, "/signup")
		return
	}
	user := models.User{First_name: data.fname, Last_name: data.lname, Username: data.uname, Email: data.mail, Password: data.password}
	DB := config.DBConn()
	result := DB.Create(&user)
	if result.Error != nil {
		c.Redirect(http.StatusMovedPermanently, "/signup")
		return
	} else {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}

}

func Loginpage(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("views/login.html"))
	tmpl.Execute(c.Writer, nil)
}

func Loginuser(c *gin.Context) {
	type data struct {
		uname    string
		password string
	}
	var userdata data
	c.Request.ParseForm()
	userdata.uname = c.Request.PostForm["username"][0]
	userdata.password = c.Request.PostForm["password"][0]
	if userdata.uname == "" || userdata.password == "" {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
	DB := config.DBConn()
	var temp_user models.User
	result := DB.First(&temp_user, "username LIKE ? AND password LIKE ?", userdata.uname, userdata.password)
	if result.Error != nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/home")
	}

}

func Homepage(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("views/home.html"))
	tmpl.Execute(c.Writer, nil)
}

func Adminloginpage(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("views/adminlogin.html"))
	tmpl.Execute(c.Writer, nil)
}
func Adminlogin(c *gin.Context) {
	c.Request.ParseForm()
	type data struct {
		uname    string
		password string
	}
	var userdata data
	var temp_user models.User
	userdata.uname = c.Request.PostForm["username"][0]
	userdata.password = c.Request.PostForm["password"][0]
	DB := config.DBConn()
	result := DB.First(&temp_user, "username LIKE ? AND password LIKE ? AND is_admin LIKE ?", userdata.uname, userdata.password, "yes")
	if result.Error != nil {
		c.Redirect(http.StatusMovedPermanently, "/adminloginpage")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/adminpanel")
	}

}

func Adminpanel(c *gin.Context) {

	var temp_user []models.User
	DB := config.DBConn()
	result := DB.Find(&temp_user)
	if result.Error != nil {
		c.Redirect(http.StatusMovedPermanently, "/adminloginpage")
	} else {
		tmpl := template.Must(template.ParseFiles("views/adminpanel.html"))
		tmpl.Execute(c.Writer, temp_user)
	}

}
func Delete(c *gin.Context) {
	c.Request.ParseForm()
	delete_id := c.Request.PostForm["id"]
	DB := config.DBConn()
	var temp_user models.User
	DB.Delete(&temp_user, delete_id)
	c.Redirect(http.StatusMovedPermanently, "/adminpanel")
}
func Search(c *gin.Context) {
	c.Request.ParseForm()
	search_uname := c.Request.PostForm["username"]
	DB := config.DBConn()
	var temp_user []models.User
	result := DB.First(&temp_user, " username LIKE ? ", search_uname)
	if result.Error != nil {
		c.Redirect(http.StatusMovedPermanently, "/adminpanel")

	} else {
		tmpl := template.Must(template.ParseFiles("views/adminpanel.html"))
		tmpl.Execute(c.Writer, temp_user)

	}

}
