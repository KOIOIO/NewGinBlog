package Model

import (
	"NewGinBlog/Utills"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var Db *gorm.DB
var err error

func InitDb() {
	Db, err = gorm.Open(Utills.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		Utills.DbUser,
		Utills.DbPassWord,
		Utills.DbHost,
		Utills.DbPort,
		Utills.DbName,
	))
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数", err)
	}
	Db.SingularTable(true)

	Db.AutoMigrate(&User{}, &Category{}, &Article{})
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(10 * time.Second)
}
