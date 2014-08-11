package models

import (
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/toolbox"
	_ "github.com/go-sql-driver/mysql"
)

type Options struct {
	Id     int
	Option string
	Value  string
}

/**
 * 获取所有的配置和值
 */
func (this *Options) GetOptions() map[string]string {
	ret := make(map[string]string, 0)
	allOptions := make([]*Options, 0, 1000)
	o.QueryTable("options").All(&allOptions)
	for key, _ := range allOptions {
		ret[allOptions[key].Option] = allOptions[key].Value
	}
	return ret
}

/**
 * 获取所有的选项key
 */
func (this *Options) GetOptionsKey() []string {
	var lists orm.ParamsList
	ret := make([]string, 0, 1000)
	o.QueryTable("options").ValuesFlat(&lists, "option")
	for _, option := range lists {
		ret = append(ret, option.(string))
	}
	return ret
}

/**
 * 入库设置
 */
func (this *Options) SetOptions(canOptions map[string]string) error {
	for option, value := range canOptions {
		_, err := o.QueryTable("options").Filter("option", option).Update(orm.Params{"value": value})
		if err != nil {
			return err
		}
	}
	return nil
}
