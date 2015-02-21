package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) Post() {
	content := this.Input().Get("content")
	tid := this.Input().Get("tid")
	err := models.AddComment(tid, content)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic/view/"+tid, 301)
}
