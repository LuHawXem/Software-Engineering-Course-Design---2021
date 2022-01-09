package service

import (
	"book-management/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken.X-CSRF-Token,Authorization,Token")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func BeforeLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.PostForm("username")
		password := c.PostForm("password")

		if name == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"errmsg": "param missing",
			})
			c.Abort()
			return
		}

		var query models.User
		models.GetGorm().Where("name = ?", name).First(&query)
		if query.Id != 0 {
			if query.Password != password {
				c.JSON(http.StatusBadRequest, gin.H{
					"errmsg": "password Error",
				})
				c.Abort()
				return
			}
			c.Set("user_id", query.Id)
			c.Set("role_id", query.RoleId)
		}

		c.Next()
	}
}

func Before() gin.HandlerFunc  {
	return func(c *gin.Context) {
		_, v := c.Get("user_id")
		if !v {
			c.JSON(http.StatusBadRequest, gin.H{
				"errmsg": "user not login",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
