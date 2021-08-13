/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 03:31:59
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 06:58:27
 * @Description:
 * @FilePath: /super_high_concurrency_system/model/buyRecord.go
 * @GitHub: https://github.com/liuyuchen777
 */

package model

import (
	"fmt"
	"log"
)

type Record struct {
	Id          int64  `gorm:"primary_key;auto_increment"`
	Username    string `gorm:"type:varchar(60); not null"`
	ProductName string `gorm:"type:varchar(60); not null"`
}

type RecordMsg struct {
	Username string
	ProductName string
}

func CreateRecord(username, productname string) error {
	if db != nil {
		record := Record{
			Username:    username,
			ProductName: productname,
		}
		result := db.Create(&record)
		fmt.Println("Insert Record: ", record.Id)
		return result.Error
	} else {
		log.Fatal("DB is not opened!")
		return nil
	}
}
