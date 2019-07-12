package model

import "time"

type Bigdata_Task_project struct {
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

func (Bigdata_Task_project) Tablename() string {
	return "bigdata_task_project"

}

//全局禁用表名负数
//db.SingularTable(true)

//db.Set("gorm:table_options","ENGINE=InnoDB")