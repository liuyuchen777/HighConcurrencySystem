/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 02:56:11
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 06:36:50
 * @Description:
 * @FilePath: /spike_system/model/product.go
 * @GitHub: https://github.com/liuyuchen777
 */
package model

import (
	"fmt"
	"log"
)

type Product struct {
	Id          int64  `gorm:"primary_key;auto_increment"`
	Name        string `gorm:"type:varchar(60); not null"` // name of product
	Amount      int32  `gorm:"type:int;not null"`          // total number
	Left        int32  `gorm:"type:int;not null"`          // left number
	Price       int32  `gorm:"type:int;not null"`
	Description string `gorm:"type:varchar(60)"` // descrption of product
}

// between redis and db
type ProductLeft struct {
	Name string
	Left int32
}

// create one product
func CreateProduct(name string, amount, left, price int32, description string) error {
	if db != nil {
		product := Product{
			Name:        name,
			Amount:      amount,
			Left:        left,
			Price:       price,
			Description: description,
		}
		result := db.Create(&product)
		if result.Error != nil {
			return result.Error
		} else {
			fmt.Println("Insert Product: ", product.Id)
			return nil
		}
	} else {
		// Future: self-define error
		log.Fatal("DB is not opened!")
		return nil
	}

}

// return all product
func GetAllProduct() ([]Product, error) {
	if db != nil {
		var products []Product

		result := db.Find(&products)

		if result.Error != nil {
			return nil, result.Error
		} else {
			return products, nil
		}
	} else {
		// Future: self-define error
		log.Fatal("DB is not opened!")
		return nil, nil
	}

}

// decrease product left
// update
func DecreaseProductLeft(username, productname string) error {
	if db != nil {
		var product Product
		db.Where("name = ?", productname).First(&product)
		product.Left--
		db.Save(&product)
		// generate a new buy record
		return CreateRecord(username, productname)
	} else {
		log.Fatal("DB is not opened!")
		return nil
	}
}

// update left
func UpdateLeft(productname string, left int32) error {
	if db != nil {
		var product Product
		db.Where("name = ?", productname).First(&product)
		product.Left = left
		db.Save(&product)
		// generate a new buy record
		return nil
	} else {
		log.Fatal("DB is not opened!")
		return nil
	}
}
