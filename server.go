/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 02:21:50
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 06:33:23
 * @Description:
 * @FilePath: /spike_system/spike_system.go
 * @GitHub: https://github.com/liuyuchen777
 */
package main

import (
	"spike_system/router"
)

const port = "9090"

func main() {
	r := router.InitEngine()

	r.Run(":" + port)
}
