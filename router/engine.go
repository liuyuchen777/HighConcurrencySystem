/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 02:28:43
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 08:36:19
 * @Description:
 * @FilePath: /spike_system/router/engine.go
 * @GitHub: https://github.com/liuyuchen777
 */
package router

import (
	"spike_system/config"
	"spike_system/controller"
	"spike_system/middleware"
	"spike_system/model"
	"spike_system/routine"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitEngine() *gin.Engine {
	r := gin.Default()

	// set config
	config, err := config.GetAppConfig()
	if err != nil {
		panic("failed to load redisService config" + err.Error())
	}

	// init redis and mysql
	model.InitDB(config)
	model.InitRedis(config)
	model.PreLoad()

	// session init
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("session", store))

	// route setting
	// user auth
	userAuth := r.Group("/api/auth")
	{
		userAuth.POST("/login", controller.Login)
		userAuth.POST("/logout", controller.Logout)
	}
	// main page
	userMain := r.Group("/api/main")
	userMain.Use(middleware.AuthMiddleware())
	{
		userMain.GET("/products", controller.GetAllProducts)
		userMain.POST("/:username/buy/:product", controller.BuyProduct)
		userMain.GET("/check/:product", controller.CheckLeft)
	}

	// run go routine between redis and mysql
	go routine.BuyRecordConsumer()
	go routine.UpdateConsumer()

	return r
}
