package controllers

import (
	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (self *LoginController) Get() {
	IsExit := self.Input().Get("exit") == "true"
	if IsExit {
		//self.Ctx.SetCookie("uname", "", -1, "/")
		//self.Ctx.SetCookie("pwd", "", -1, "/")
		self.SetSession("uname","uname")
		self.SetSession("pwd","pwd")
		self.Redirect("/", 301)
		return
	}
	self.TplNames = "login.html"
}

func (self *LoginController) Post() {
	uname := self.Input().Get("uname")
	pwd := self.Input().Get("pwd")
//	autoLogin := self.Input().Get("autoLogin") == "on"
	if uname == beego.AppConfig.String("uname") && pwd == beego.AppConfig.String("pwd") {
	/*	maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}*/
		//self.Ctx.SetCookie("uname", uname, maxAge, "/")
		//self.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		self.SetSession("uname",uname)
		self.SetSession("pwd",pwd)
	}
	self.Redirect("/", 301)
	return
}

/*func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value
	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value
	return uname == beego.AppConfig.String("uname") && pwd == beego.AppConfig.String("pwd")
}
*/
