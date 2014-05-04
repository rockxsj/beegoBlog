package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/**
 * 新增文章
 */
func AddOne(title string, content string) bool {
	o := orm.NewOrm()

	var posts = new(Posts)
	posts.Title = title
	posts.Content = content
	posts.Add_time = time.Now()
	_, err := o.Insert(posts)
	if err == nil {
		return true
	} else {
		return false
	}
}

/**
 * 删除文章
 */
func Del(id int) {
	o := orm.NewOrm()

	o.Delete(&Posts{Id: id})
}

/**
 * 获取一个记录，供更新
 */
func GetOne(id int) *Posts {
	o := orm.NewOrm()

	posts := new(Posts)

	posts.Id = id
	o.Read(posts)
	return posts
}

/**
 * 获取相邻记录
 *
 * @param int 当前id
 * @param bool 之前还是之后
 */
func GetPreOrNext(id int, next bool) (error, *Posts) {
	o := orm.NewOrm()
	posts := new(Posts)

	var fiter, order string
	if next {
		fiter = "Id__gt"
		order = "Id"
	} else {
		fiter = "Id__lt"
		order = "-Id"
	}
	preOrNext := new(Posts)
	err := o.QueryTable(posts).Filter(fiter, id).OrderBy(order).Limit(1).One(preOrNext)
	if err == orm.ErrNoRows {
		return orm.ErrNoRows, preOrNext
	} else {
		return nil, preOrNext
	}
}

/**
 * 更新一条记录
 */
func DoUpdateOne(id int, title string, content string) bool {
	o := orm.NewOrm()

	posts := Posts{Id: id}
	if o.Read(&posts) == nil {
		posts.Title = title
		posts.Content = content
		posts.Mod_time = time.Now()

		if _, err := o.Update(&posts); err == nil {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
