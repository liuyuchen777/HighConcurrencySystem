/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 02:57:08
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 09:09:05
 * @Description:
 * @FilePath: /spike_system/middleware/authMiddleware.go
 * @GitHub: https://github.com/liuyuchen777
 */

package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check session
		session := sessions.Default(c)
		if cookie, err := c.Request.Cookie("session_id"); err == nil {
			// have cookie
			value := cookie.Value
			fmt.Println("Cookie Value: ", value)
			v := session.Get(value)
			fmt.Println("value of v: ", v)
			if v == nil {
				// no session record
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Wrong Cookie Unauthorized",
				})
				c.Abort()
				return
			} else {
				c.Next()
				return
			}
		} else {
			// no cookie
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "No Cookie Unauthorized",
			})
			c.Abort()
			return
		}
	}
}
