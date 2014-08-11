package models

import (
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/toolbox"
	_ "github.com/go-sql-driver/mysql"
)

type Cates struct {
	Id   int64
	Pid  int64
	Name string
}

func (this *Cates) GetCates() []*Cates {
	var cates []*Cates
	o.QueryTable("cates").All(&cates)
	return cates
}

func (this *Cates) AddCate(name string) error {
	this.Name = name
	_, err := o.Insert(this)
	return err
}

func (this *Cates) DelCate(id int64) error {
	_, err := o.Delete(&Cates{Id: id})
	o.QueryTable("posts").Filter("Cid", id).Update(orm.Params{
		"Cid": 0,
	})
	return err
}

func (this *Cates) RenameCate(id int64, newName string) error {
	cate := Cates{Id: id}
	cate.Name = newName
	_, err := o.Update(&cate)
	return err
}
