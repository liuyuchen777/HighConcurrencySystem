/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 05:36:47
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 06:21:32
 * @Description:
 * @FilePath: /super_high_concurrency_system/test/DB_test.go
 * @GitHub: https://github.com/liuyuchen777
 */
package test

import (
	"fmt"
	"super_high_concurrency_system/config"
	"super_high_concurrency_system/model"

	"testing"
)

func TestDB(t *testing.T) {
	config, err := config.GetAppConfig()
	if err != nil {
		panic("failed to load redisService config" + err.Error())
	}

	// init redis and mysql
	model.InitDB(config)
	// create some user
	model.CreateUser("LiuYuchen", "123456")
	model.CreateUser("ZhuJunYi", "980628")
	model.CreateUser("XiangZihao", "774201")
	// verify password
	user1 := "LiuYuchen"
	fmt.Println(user1, model.FindUser(user1))
	user2 := "ZhuJunYi"
	fmt.Println(user2, model.FindUser(user2))
	// add some product
	model.CreateProduct("IPhone12", 5, 5, 100, "A really good smart phone.")
	model.CreateProduct("MacM1", 3, 3, 200, "A really good computer.")
	// get all
	res, _ := model.GetAllProduct()
	for _, v := range res {
		fmt.Println(v.Name, v.Left)
	}
	// decrease amount left
	model.DecreaseProductLeft("LiuYuchen", "MacM1")
	res, _ = model.GetAllProduct()
	for _, v := range res {
		fmt.Println(v.Name, v.Left)
	}
}
