/*
*
@author: Azir
@desc:
@date: 6/19/24
*
*/
package common

import (
	"backend/model"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	user := "root"
	password := "12345678"
	host := "localhost"
	port := "3306"
	database := "blog"
	charset := "utf8mb4"
	loc := "Asia/Shanghai"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		user,
		password,
		host,
		port,
		database,
		charset,
		// url.QueryEscape 对字符串进行转义，使其可以安全地用在 URL 中
		url.QueryEscape(loc))
	// 连接数据库
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	// 迁移数据表
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

// 数据库信息获取
func GetDB() *gorm.DB {
	return DB
}
