package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {

	this.Data["IsTopic"] = true
	var err error
	this.Data["Topics"], err = models.GetAllTopics("")
	if err != nil {
		beego.Error(err)
	}
	this.TplNames = "topic.html"
}

func (this *TopicController) Post() {
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	cid := this.Input().Get("cid")
	id := this.Input().Get("id")
	//this.Ctx.WriteString(id)
	var err error
	if len(id) == 0 {
		err = models.AddTopic(title, content, cid)
	} else {
		err = models.ModifyTopic(id, title, content)
	}
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
	return
}

func (this *TopicController) Add() {
	//this.Ctx.WriteString("add")
	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.TplNames = "topic_add.html"
}

func (this *TopicController) Modify() {
	id := this.Ctx.Input.Params["0"]
	if len(id) == 0 {
		this.Redirect("/topic", 301)
		return
	}
	//this.Ctx.WriteString(id)
	var err error
	this.Data["Topic"], err = models.GetTopicById(id)
	if err != nil {
		beego.Error(err)
	}
	this.TplNames = "topic_modify.html"
}

func (this *TopicController) Delete() {
	id := this.Ctx.Input.Params["0"]
	if len(id) == 0 {
		this.Redirect("/topic", 301)
		return
	}
	err := models.DelTopic(id)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 301)
}

func (this *TopicController) View() {
	tid := this.Ctx.Input.Params["0"]
	//this.Ctx.WriteString(tid)
	var err error
	this.Data["Topic"], err = models.GetTopicById(tid)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Comments"], err = models.GetAllComments(tid)
	if err != nil {
		beego.Error(err)
	}
	this.TplNames = "view.html"
}
