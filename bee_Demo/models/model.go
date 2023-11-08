package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //导入驱动包
)

/*
*
# Mysql Config
db_name = go_test
db_user = root
db_pass = 123456
db_url =127.0.0.1
port = 3306
*/
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	user := beego.AppConfig.String("db_user")
	pwd := beego.AppConfig.String("db_pass")
	host := beego.AppConfig.String("db_url")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("db_name")
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	err := orm.RegisterDataBase("default", "mysql", dbConn, 30)

	if err != nil {
		logs.Info("连接数据库出错")
		return
	}
	logs.Info("连接数据库成功")
	//register model注册表
	orm.RegisterModel(
		new(Permission),
		new(Admin),
		new(User))
	orm.RunSyncdb("default", false, true)
}

//==========================================数据库表结构===============================================

// 管理员权限等级及级别名称
type Permission struct {
	Id    int      //权限登记id
	Level string   `json:"level" orm:"size(30)"` //权限级别
	Name  string   `json:"name"  orm:"size(20)"` //权限名称
	Admin []*Admin `orm:"rel(m2m)"`              //orm映射 一个权限可以被多个管理员所拥有
}

/**
*管理员表
 */
type Admin struct {
	Id         int    `json:"id"`                        //管理员编号id
	UserName   string `json:"user_name"  orm:"size(12)"` //管理员用户名
	CreateTime string `json:"create_time"`               //记录添加时间
	Status     int    `json:"status"`                    //管理员状态
	Avatar     string `json:"avatar" orm:"size(50)"`     //管理员头像
	Pwd        string `json:"pwd"`                       //管理员密码
	//Permission []*Permission `orm:"reverse(many)"`              //一个管理员可以由多种权限
	//City       *City         `orm:"rel(fk)"`                    //orm映射 管理员所在城市

}

/**
 * 用户信息表
 */
type User struct {
	Id          int    `json:"id"`           //用户的编号id
	UserName    string `json:"username"`     //用户名称
	RegisteTime string `json:"registe_time"` //用户注册时间
	Mobile      string `json:"mobile"`       //用户的移动手机号
	IsActive    int    `json:"is_active"`    //用户是否激活
	Balance     int    `json:"balance"`      //用户的账户余额（简单起见，使用int类型）
	Avatar      string `json:"avatar"`       //用户的头像
	//City        *City        `orm:"rel(fk)"`       //orm映射 用户所在城市 一对一关系 一个用户能有一个城市地区
	//UserOrder   []*UserOrder `orm:"reverse(many)"` //orm映射 用户的订单 一个用户可以有多张订单，设置一对多的关系
	Pwd     string `json:"password"` //用户的账户密码
	DelFlag int    `json:"del_flag"` //是否被删除的标志字段 软删除
}
