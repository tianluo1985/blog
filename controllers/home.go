package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	//uname := c.Ctx.GetCookie("uname")
	//pwd := c.Ctx.GetCookie("pwd")

	//c.Data["uname"] = uname
	//c.Data["pwd"] = pwd
	cid := c.Input().Get("cid")

	c.Data["IsHome"] = true

	c.Data["IsLogin"] = c.checkAccount()
	var err error
	c.Data["Topics"], err = models.GetAllTopics(cid)
	c.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	c.TplNames = "home.html"
}

func (this *HomeController) checkAccount() bool{
	uname:=this.GetSession("uname")
	if uname==nil{
		return false
	}
	pwd:=this.GetSession("pwd")
	if pwd==nil{
		return false
	}
	return uname == beego.AppConfig.String("uname") && pwd == beego.AppConfig.String("pwd")
}