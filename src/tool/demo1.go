package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"math/rand"
	"time"
)


var db *gorm.DB

type Project struct {
	ID int `gorm:"primary_key"`
	Name string `gorm:"type:varchar(64);not null"`
	Title string `gorm:"type:varchar(64);unique;not null"`
	Runtime string `gorm:"type:varchar(64);"`
	Type int `gorm:"default:0;"`
	LastTime time.Time
	Dingmachince string `gorm:"type:varchar(255);"`
	Filename string `gorm:"type:varchar(255);"`
	Ossaddress string `gorm:"size:512"`
	UpdateTime time.Time
	CreateTime time.Time
	IsDelete int `gorm:"default:0;not null"`
	Desc string `gorm:"type:varchar(255);"`
	Codeinfo string `gorm:"type:longText"`
	Ossfilename string `gorm:"type:varchar(64);"`
}

func (t * Project) Tablename() string {
	fmt.Println("设备表名")
	return "testtprojects"

}
func GetRandomString(l int) string  {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes:=[]byte(str)
	result:=[]byte{}
	r:=rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0;i<l;i++{
		result=append(result,bytes[r.Intn(len(bytes))])
	}
	//fmt.Println(result)
	return string(result)

}
//func (project *Project)  beforeCreateCallback1(scope gorm.Scope) error{
//	fmt.Println("插入的时候可以设置回调函数，在回调函数中做一系列的处理")
//	fmt.Println(uuid.NewV1())
//	scope.SetColumn("ID",uuid.NewV1())
//	return nil
//
//}
//func beforeCreate(scope gorm.Scope) error{
//	fmt.Println("插入的时候可以设置回调函数，在回调函数中做一系列的处理")
//	//fmt.Println(uuid.NewV1())
//	//scope.SetColumn("ID",uuid.NewV1())
//	return nil
//
//}
//func beforeCreate(scope *gorm.Scope) {
//	fmt.Println("插入的时候可以设置回调函数，在回调函数中做一系列的处理")
//
//}
//
//func (t *Project) AfterCreate() (err error) {
//	fmt.Println("回调示例")
//	return nil
//
//}

//
//func (u *Project) BeforeCreate(scope *gorm.Scope) (err error) {
//	fmt.Println("回调示例,BeforCreate()")
//	fmt.Println(uuid.NewV1())
//	scope.SetColumn("ID",uuid.NewV1())
//	return
//}
//func (u *Project) Create(scope *gorm.Scope) (err error) {
//	fmt.Println("回调示例,Create()")
//	return
//}
//
//func (u *Project) BeforeUpdate(scope *gorm.Scope) (err error) {
//	now:=gorm.NowFunc()
//	fmt.Println("回调示例,Before()")
//	if _, ok := scope.Get("gorm:update_time"); !ok {
//		scope.SetColumn("UpdateTime", now)
//	}
//	return
//}
//
//func (u *Project) AfterCreate(scope *gorm.Scope) (err error) {
//	fmt.Println("回调示例,AfterCreate()")
//	if !scope.HasError(){
//		now:=gorm.NowFunc()
//		fmt.Println(now)
//		scope.DB().Model(u).Update("UpdateTime", now)
//		scope.DB().Model(u).Update("createTime", now)
//
//
//		if createdAtField, ok := scope.FieldByName("CreateTime"); ok {
//			if createdAtField.IsBlank {
//				createdAtField.Set(now)
//			}
//		}
//
//		if updatedAtField, ok := scope.FieldByName("UpdateTime"); ok {
//				updatedAtField.Set(now)
//		}
//	}
//
//
//	return
//}

func (t *Project) updateTimeStampForCreateCallback(scope *gorm.Scope) {
	fmt.Println("updateTimeStampForCreateCallback回调函数")
	if !scope.HasError() {
		now := gorm.NowFunc()

		if createdAtField, ok := scope.FieldByName("CreatedAt"); ok {
			if createdAtField.IsBlank {
				createdAtField.Set(now)
			}
		}

		if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
			if updatedAtField.IsBlank {
				updatedAtField.Set(now)
			}
		}
	}
}


func (t *Project) AfterCreate( scope * gorm.Scope)  {
	fmt.Println("AfterCreate回调函数")
	return

}

func (t *Project) BeforeCreate( scope * gorm.Scope)  {
	fmt.Println("BeforeCreate回调函数")
	if !scope.HasError() {
		now := gorm.NowFunc()

		if createdAtField, ok := scope.FieldByName("CreateTime"); ok {
			if createdAtField.IsBlank {
				createdAtField.Set(now)
			}
		}

		if updatedAtField, ok := scope.FieldByName("UpdateTime"); ok {
			if updatedAtField.IsBlank {
				updatedAtField.Set(now)
			}
		}
	}

	return

}
func (t *Project) AfterSave( scope * gorm.Scope)  {
	fmt.Println("AfterSave回调函数")
	return

}
func (t *Project) BeforeSave( scope * gorm.Scope)  {
	fmt.Println("BeforeSave回调函数")
	return

}

func (t *Project) BeforeUpdate( scope * gorm.Scope)  {
	fmt.Println("BeforeUpdate回调函数")
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdateTime", gorm.NowFunc())
	}
	return

}

func (t *Project) AfterUpdate( scope * gorm.Scope)  {
	fmt.Println("AfterUpdate回调函数")
}

func createNew(db *gorm.DB)  {
	title:=GetRandomString(6)
	filename:=GetRandomString(5)+".py"
	ding:=GetRandomString(10)
	name:=GetRandomString(6)
	var project Project
	project=Project{Name:name,Title:title,Runtime:"11 * * * *",Dingmachince:ding,Filename:filename}
	//a:=db.NewRecord(project)
	//fmt.Println(a)
	db.Create(&project)

}


func Init()  (* gorm.DB){
	fmt.Println("init")
	db, err := gorm.Open("mysql", "aitest:7eXYRaLWDlA3WUmo@tcp(rm-2zevlk47ul0ovuci80o.mysql.rds.aliyuncs.com)/aitest?charset=utf8&parseTime=true")
	if err!=nil{
		fmt.Println(err)
		panic("连接数据库失败")
	}
	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(2000)
	db.DB().SetMaxOpenConns(1000)

	gorm.DefaultTableNameHandler= func(db *gorm.DB, defaultTableName string) string {
		return "test_"+defaultTableName

	}
	db.Callback().Create().Register("gorm:before_create1", beforeCreate1)
	db.Set("grom:table_options","ENGINE=InnoDB")
	fmt.Println("数据库初始化完成")
	return db


}

func updateNew(db *gorm.DB)  {
	var project Project
	db.Where("name=?","ddlo65").First(&project)
	fmt.Println(project)

	db.Model(&project).Where("name=?","ddlo65").Update("name","gormtest")

}
func main()  {
	db=Init()
	print("main")
	defer db.Close()
	//createNew(db)
	updateNew(db)
}

func main1()  {
	defer db.Close()
	db.AutoMigrate(&Project{})

	title:=GetRandomString(6)
	filename:=GetRandomString(5)+".py"
	ding:=GetRandomString(10)
	name:=GetRandomString(6)
	var project Project
	project=Project{Name:name,Title:title,Runtime:"11 * * * *",Dingmachince:ding,Filename:filename}
	a:=db.NewRecord(project)
	fmt.Println(a)
	db.Create(&project)
	b:=db.NewRecord(project)
	fmt.Println(b)

	db.First(&project)
	fmt.Println(project)
	project.Name="cyyy"
	db.Save(&project)
	find_all(db)

	//db.Model(&Project{}).Where("runtime=?","11 * * * *").Update("name","曹亚运")
	var pro []Project
	fmt.Println("search")
	db.Raw("SELECT * FROM test_project WHERE name = ?", "曹亚运").Scan(&pro)
	fmt.Println(pro[0])

}
func find_all(db *gorm.DB) {
	var project []Project
	db.First(&project)
	fmt.Println(project)


}
func beforeCreate1(scope *gorm.Scope) {
	fmt.Println("自定义回调示例回调,BeforCreate1()")
	return

}





//全局禁用表名负数
//db.SingularTable(true)

//db.Set("gorm:table_options","ENGINE=InnoDB")
