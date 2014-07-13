package models

import (
	//"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/toolbox"
	_ "github.com/go-sql-driver/mysql"
)

type Links struct {
	Id   int64
	Name string
	Link string
}

/**
 * 添加友情链接
 */
func (this *Links) AddLink(name, link string) {
	this.Name = name
	this.Link = link
	o.Insert(this)
}

/**
 * 删除友情链接
 */
func (this *Links) DelLink(id int64) error {
	_, err := o.Delete(&Links{Id: id})
	return err
}

/**
 * 获取全部友情链接
 */
func (this *Links) GetLinks() []*Links {
	var allLinks []*Links
	o.QueryTable("links").All(&allLinks)
	return allLinks
}
