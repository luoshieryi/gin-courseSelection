package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			//config.Mysql.Username,
			//config.Mysql.Password,
			//config.Mysql.Host,
			//config.Mysql.Port,
			//config.Mysql.Database,
			"root",
			"bytedancecamp",
			"180.184.74.141",
			3306,
			"courseSelection",
		),
	)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	//显示SQL
	DB.LogMode(true)

	//设置连接池
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	//数据库迁移
	migrate()
}

func migrate() {
	DB.AutoMigrate(&Member{})
	DB.AutoMigrate(&Session{})
	DB.AutoMigrate(&Course{})
	DB.AutoMigrate(&StudentCourse{})
}

func Close() {
	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {
			panic(err)
		}
	}(DB)
}

func init() {
	Init()
}
