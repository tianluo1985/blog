package models

import (
	//"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/mattn/go-sqlite3"
	//"os"
	//"path"
	"strconv"
	"time"
)

/*const (
	_DB_NAME        = "/data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)*/

type Comment struct {
	Id      int64
	TId     int64
	Created time.Time
	Content string `orm:"size(2000)"`
}

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

/*type Category struct {
	Id         int64
	Title      string
	TopicCount int64
}*/

type Topic struct {
	Id              int64
	UId             int64
	CId             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          int64
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	/*	if !com.IsExist(_DB_NAME) {
			os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
			os.Create(_DB_NAME)
		}

		orm.RegisterModel(new(Category), new(Topic))
		orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
		orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)*/

	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/beeblog?charset=utf8", 10)
}

func AddTopic(title, content, cid string) error {
	o := orm.NewOrm()
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}
	topic := &Topic{Title: title, Content: content, CId: cidNum, Created: time.Now(), Updated: time.Now(), ReplyTime: time.Now()}
	_, err = o.Insert(topic)
	return err
}

func AddComment(tid, content string) error {
	o := orm.NewOrm()
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	comment := &Comment{TId: tidNum, Content: content, Created: time.Now()}
	_, err = o.Insert(comment)
	return err
}

func GetTopicById(id string) (*Topic, error) {
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	topic := &Topic{Id: idNum}
	err = qs.Filter("Id", idNum).One(topic)
	return topic, err
}

func ModifyTopic(id string, title string, content string) error {
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: idNum}
	err = o.Read(topic)
	if err == nil {
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		_, err = o.Update(topic)
	}
	return err
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	//o.Using("default")
	cate := &Category{Title: name, Created: time.Now(), TopicTime: time.Now()}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	//o.Insert(cate)
	//var err error
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func GetAllComments(tid string) ([]*Comment, error) {
	comments := make([]*Comment, 0)
	o := orm.NewOrm()
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	qs := o.QueryTable("comment").Filter("TId", tidNum)
	_, err = qs.All(&comments)
	return comments, err
}
func GetAllTopics(cid string) ([]*Topic, error) {
	topics := make([]*Topic, 0)
	o := orm.NewOrm()
	var err error
	if len(cid) > 0 {
		cidNum, err := strconv.ParseInt(cid, 10, 64)
		if err != nil {
			return nil, err
		}
		qs := o.QueryTable("topic")
		qs = qs.Filter("CId", cidNum)
		_, err = qs.All(&topics)
	} else {
		qs := o.QueryTable("topic")
		_, err = qs.All(&topics)
	}
	return topics, err
}

func GetAllCategories() ([]*Category, error) {
	cates := make([]*Category, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func DelTopic(id string) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	topic := &Topic{Id: tid}
	o := orm.NewOrm()
	_, err = o.Delete(topic)
	return err
}

func DelCategory(str_id string) error {
	cid, err := strconv.ParseInt(str_id, 10, 64)
	if err != nil {
		return err
	}
	cate := &Category{Id: cid}
	o := orm.NewOrm()
	_, err = o.Delete(cate)
	return err
}
