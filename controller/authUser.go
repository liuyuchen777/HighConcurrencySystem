/*
 * @Author: Liu Yuchen
 * @Date: 2021-05-08 02:48:43
 * @LastEditors: Liu Yuchen
 * @LastEditTime: 2021-05-08 09:03:00
 * @Description:
 * @FilePath: /spike_system/controller/authUser.go
 * @GitHub: https://github.com/liuyuchen777
 */
package controller

import (
	"fmt"
	"log"
	"net/http"
	"spike_system/model"
	"spike_system/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	// login function
	var user LoginUser
	c.BindJSON(&user)
	var info string
	temp := model.FindUser(user.Username)

	if temp == "" {
		// no such user
		info = "No Such User!"
	} else if temp != user.Password {
		// password incorrect
		info = "Password Incorrect"
	} else {
		// login successful
		msg := utils.GetSHA256HashCode(user.Username)
		// generate shr256id
		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    msg,
			Path:     "/",
			HttpOnly: true,
		}
		info = "Login Success!"
		http.SetCookie(c.Writer, cookie)
		session := sessions.Default(c)
		session.Set(msg, user.Username)
		session.Save()
	}
	// return info
	c.JSON(http.StatusOK, gin.H{
		"info": info,
	})
}

func Logout(c *gin.Context) {
	// logout function
	session := sessions.Default(c)
	if cookie, err := c.Request.Cookie("session_id"); err == nil {
		// have cookie
		value := cookie.Value
		fmt.Println(value)
		session.Delete(value)
		session.Save()
		c.JSON(http.StatusOK, gin.H{
			"info": "Logout Successfully!",
		})
	} else {
		// something strange happened, no session but login
		log.Fatal("something strange happened!")
	}
}
