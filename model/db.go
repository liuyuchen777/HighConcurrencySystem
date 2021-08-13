/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 03:46:42
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 08:37:06
 * @Description:
 * @FilePath: /super_high_concurrency_system/model/db.go
 * @GitHub: https://github.com/liuyuchen777
 */
package model

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fmt"
	"super_high_concurrency_system/config"
)

var db *gorm.DB

func AddUser() {
	CreateUser("LiuYuchen", "123456")
	CreateUser("ZhuJunYi", "980628")
	CreateUser("XiangZihao", "774201")
}

func AddProduct() {
	CreateProduct("IPhone12", 5, 5, 100, "A really good smart phone.")
	CreateProduct("MacM1", 3, 3, 200, "A really good computer.")
	// get all
	res, _ := GetAllProduct()
	for _, v := range res {
		fmt.Println(v.Name, v.Left)
	}
}

func InitDB(config config.AppConfig) {
	var err error
	// connect database
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.App.Database.User, config.App.Database.Password, config.App.Database.Address,
		config.App.Database.DbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Can't Connect Database!")
	}
	// migeate database
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Record{})

	// drop all table
	db.Where("1 = 1").Delete(&User{})
	db.Where("1 = 1").Delete(&Product{})
	db.Where("1 = 1").Delete(&Record{})

	// load some data
	AddUser()
	AddProduct()
}
