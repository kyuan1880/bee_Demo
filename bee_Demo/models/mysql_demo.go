package models

//
//import (
//	"fmt"
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/orm"
//	_ "github.com/go-sql-driver/mysql" // import your used driver
//)
//
//type Test_User struct {
//	Id   int
//	Name string
//}
//
////func init() {
////
////	// set default database
////	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8", 30)
////
////	// register model
////	//orm.RegisterModel(new(Test_User))
////
////	// create table
////	orm.RunSyncdb("default", false, true)
////}
//
//func MysqlTest() {
//	post := beego.AppConfig.String("post")
//	fmt.Println("post:", post)
//	fmt.Println("mysql conn success!")
//	o := orm.NewOrm()
//	user := Test_User{Name: "lisha"}
//	id, err := o.Insert(&user)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Printf("ID: %d, ERR: %v\n", id, err)
//}
