/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 02:54:48
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 08:33:00
 * @Description:
 * @FilePath: /super_high_concurrency_system/controller/getProduct.go
 * @GitHub: https://github.com/liuyuchen777
 */
package controller

import (
	"net/http"
	"super_high_concurrency_system/model"
	"super_high_concurrency_system/routine"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	products, err := model.GetAllProduct()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"products": nil,
			"info":     "Can't Get Product!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"products": products,
			"info":     "Get All Products Successful!",
		})
	}
}

func BuyProduct(c *gin.Context) {
	productname := c.Param("product")
	username := c.Param("username")

	left, info := model.BuyProductRedis(productname)
	if info == "Buy Successfully!" {
		// buy successfully
		// use channel to transmit to db
		routine.BuyRecordChannel <- model.RecordMsg{
			Username:    username,
			ProductName: productname,
		}
		routine.UpdateLeftChannel <- model.ProductLeft{
			Name: productname,
			Left: left,
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"info": info,
	})
}

func CheckLeft(c *gin.Context) {
	productname := c.Param("product")

	val := model.GetProductLeftRedis(productname)

	if val > 0 {
		c.JSON(http.StatusOK, gin.H{
			"left": val,
			"info": "There are products left!",
		})
	} else if val == 0 {
		c.JSON(http.StatusOK, gin.H{
			"left": 0,
			"info": "All sold!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"left": -1,
			"info": "No Such Product!",
		})
	}
}
