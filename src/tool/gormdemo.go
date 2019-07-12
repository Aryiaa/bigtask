package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name string
}

type Profile struct {
	gorm.Model
	UserID int
	User User
	Name string
}
type Base  struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
type Product struct {
	Base
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "aitest:7eXYRaLWDlA3WUmo@tcp(rm-2zevlk47ul0ovuci80o.mysql.rds.aliyuncs.com)/aitest?charset=utf8&parseTime=true")

	if err != nil {
		panic("连接数据库错误")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})


	db.Create(&Product{Code:"L1212",Price:1000})
	var product Product

	a:=db.First(&product,1)
	fmt.Println(a.Value)
	b:=db.First(&product,"code=?","L1212")
	fmt.Println(b.Value)
	db.Model(&product).Update("Price",20)
	fmt.Println(product)
	fmt.Println(&product)
	db.Delete(&product)




}
