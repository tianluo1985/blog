package routers

import (
	"beeblog/controllers"
	"beeblog/models"
	"github.com/astaxie/beego"
)

func init() {
	models.RegisterDB()
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/comment", &controllers.CommentController{})
	beego.AutoRouter(&controllers.TopicController{})
}
