/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 02:56:02
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 09:04:34
 * @Description:
 * @FilePath: /super_high_concurrency_system/model/user.go
 * @GitHub: https://github.com/liuyuchen777
 */
package model

import (
	"fmt"
	"log"
)

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id       int    `gorm:"primary_key;auto_increment"`
	Username string `gorm:"type:varchar(20)"`
	Password string `gorm:"type:varchar(32)"`
}

// add one user to database
func CreateUser(username, password string) error {
	if db != nil {
		user := User{Username: username, Password: password}
		result := db.Create(&user)
		if result.Error != nil {
			return result.Error
		} else {
			fmt.Println("Insert User: ", user.Id)
			return nil
		}
	} else {
		// Future: self-define error
		log.Fatal("DB is not opened!")
		return nil
	}
}

// find username and password
// return password if found
func FindUser(username string) string {
	if db != nil {
		var user User
		db.Where("Username = ?", username).First(&user)
		if user.Password != "" {
			return user.Password
		} else {
			fmt.Println("can't find such user")
			return ""
		}
	} else {
		// Future: self-define error
		log.Fatal("DB is not opened!")
		return ""
	}
}
