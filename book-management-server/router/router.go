package router

import (
	"book-management/service"
	"github.com/gin-gonic/gin"
)

func init()  {
	r := gin.Default()
	r.Use(service.Cors())
	r.Use(service.BeforeLogin())
	r.POST("/login", service.Login)
	r.POST("/booklist", service.Before(), service.GetData)
	r.POST("/borrow", service.Before(), service.Borrow)
	r.POST("/manage", service.Before(), service.Manage)
	r.Run(":8000")
}