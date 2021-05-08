/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 03:46:45
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 08:58:26
 * @Description:
 * @FilePath: /spike_system/model/redis.go
 * @GitHub: https://github.com/liuyuchen777
 */
package model

import (
	"fmt"
	"log"
	"spike_system/config"
	"strconv"

	"github.com/go-redis/redis/v7"
)

var client *redis.Client

func InitRedis(config config.AppConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     config.App.Redis.Address,
		Password: config.App.Redis.Password, // It's ok if password is "".
		DB:       0,                         // use default DB
	})

	if _, err := FlushAll(); err != nil {
		println("Error when flushAll. " + err.Error())
	}
}

// clean all the key
func FlushAll() (string, error) {
	return client.FlushAll().Result()
}

// direct interact with redis, abandon lua script
// preload (productname, left) to redis
func PreLoad() {
	products, err := GetAllProduct()
	if err != nil {
		log.Fatal("Error Get All Products: ", err)
	}

	for _, product := range products {
		err := CacheProductRedis(product)
		if err != nil {
			log.Fatal("Cache Product in Redis Fail: ", err)
		}
	}

	fmt.Print("---Set redis keys of coupons success.---")
}

// cache product in redis
func CacheProductRedis(product Product) error {
	if client != nil {
		return client.Set(product.Name, product.Left, 0).Err()
	} else {
		log.Fatal("Redis is not opened!")
		return nil
	}
}

// check product left in redis
// return int code
// -1 no such product
// 0 product no left
// >0 number of product
func GetProductLeftRedis(productname string) int32 {
	if client != nil {
		val, err := client.Get(productname).Result()
		if err != nil {
			// no such key
			return -1
		}
		val_num, _ := strconv.ParseInt(val, 10, 32)
		if val_num == 0 {
			// no product left
			return 0
		} else {
			// left product
			return int32(val_num)
		}
	} else {
		log.Fatal("Redis is not opened!")
		return -3
	}
}

// user buy one product
func BuyProductRedis(productname string) (int32, string) {
	val := GetProductLeftRedis(productname)
	if val == -1 {
		return -1, "No Such Product!"
	} else if val == 0 {
		return -1, "No Product Left!"
	} else {
		// have product left
		val--
		err := client.Set(productname, val, 0).Err()
		if err != nil {
			log.Fatal("Update Key Error: ", err)
		}
		return val, "Buy Successfully!"
	}
}
