package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"sync"
)

var db *gorm.DB
var initOnce sync.Once

//sync.Once.Do(f func())是一个挺有趣的东西,能保证once只执行一次，无论你是否更换once.Do(xx)这里的方法,这个sync.Once块只会执行一次

func Init() {
	initOnce.Do(func() {
		initDB()
	})

}
func initDB() {
	fmt.Println("数据库初始化开始")
	fmt.Println(viper.GetString("mysql.master"))
	conn, err := gorm.Open(viper.GetString("mysql.driver"), viper.GetString("mysql.master"))

	//conn,err:=gorm.Open("mysql",`aitest:7eXYRaLWDlA3WUmo@tcp(rm-2zevlk47ul0ovuci80o.mysql.rds.aliyuncs.com)/aitest?charset=utf8&parseTime=true`)

	if err != nil {
		fmt.Println("数据库连接错误")
		fmt.Println(err)
	}
	conn.DB().SetMaxIdleConns(viper.GetInt("mysql.maxIdle"))
	conn.DB().SetMaxIdleConns(viper.GetInt("mysql.maxOpen"))
	conn.LogMode(viper.GetBool("mysql.LogMode"))
	db = conn
	fmt.Println("数据库初始化完毕")
}
