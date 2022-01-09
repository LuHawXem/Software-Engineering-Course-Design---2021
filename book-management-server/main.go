package main

import (
	_ "book-management/config"
	"book-management/models"
	_ "book-management/router"
)

func main()  {
	defer models.GetGorm().Close()
}