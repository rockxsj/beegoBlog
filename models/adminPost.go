package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/**
 * 新增文章
 */
func (this *Posts) AddOne(title string, content string, cid int64) bool {
	this.Title = title
	this.Content = content
	this.Cid = cid
	this.Add_time = time.Now()
	_, err := o.Insert(this)
	if err == nil {
		return true
	} else {
		return false
	}
}

/**
 * 删除文章
 */
func (this *Posts) Del(id int64) {
	this.Id = id
	o.Delete(this)
}

/**
 * 更新一条记录
 */
func (this *Posts) DoUpdateOne(id int64, title string, content string, cid int64) bool {
	this.Id = id
	if o.Read(this) == nil {
		this.Title = title
		this.Content = content
		this.Cid = cid
		this.Mod_time = time.Now()

		if _, err := o.Update(this); err == nil {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
