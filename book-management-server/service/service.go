package service

import (
	"book-management/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {
	_, v := c.Get("user_id")
	if !v {
		name := c.PostForm("username")
		password := c.PostForm("password")
		var user = models.User{
			Name:     name,
			Password: password,
			RoleId:   0,
		}
		models.GetGorm().Create(&user)
	}

	roleId, v := c.Get("role_id")
	if !v {
		roleId = 0
	}

	c.JSON(http.StatusOK, gin.H{
		"errmsg": "request OK",
		"role_id": roleId,
	})
}

func GetData(c *gin.Context) {
	bookNum := c.PostForm("book_num")
	bookName := c.PostForm("book_name")
	author := c.PostForm("author")
	publisher := c.PostForm("publisher")
	userId, _ :=  c.Get("user_id")
	if bookNum == "" {
		bookNum = "%"
	}
	if bookName == "" {
		bookName = "%"
	}
	if author == "" {
		author = "%"
	}
	if publisher == "" {
		publisher = "%"
	}

	var data []*models.BookDetail
	models.GetGorm().Where("book_num like ? AND book_name like ? AND author like ? AND publisher like ?", bookNum, bookName, author, publisher).Order("id asc").Find(&data)
	var borrow []*models.Borrow
	models.GetGorm().Where("user_id = ? AND count > 0", userId).Order("book_id asc").Find(&borrow)

	c.JSON(http.StatusOK, gin.H{
		"errmsg": "request OK",
		"data":   gin.H{
			"booklist": data,
			"borrow": borrow,
		},
	})
	return
}

func Borrow(c *gin.Context) {
	userId, _:= c.Get("user_id")
	bookId, _ := strconv.Atoi(c.PostForm("book_id"))
	count, _ := strconv.Atoi(c.PostForm("borrow"))

	var bookDetail models.BookDetail
	models.GetGorm().Where("id = ?", bookId).First(&bookDetail)
	if bookDetail.Lend == 0 && count < 0 {
		c.JSON(http.StatusOK, gin.H{
			"errmsg": "Error, not book lend",
		})
		return
	}
	if bookDetail.Collection == 0 && count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"errmsg": "Error, not book leave",
		})
		return
	}
	bookDetail.Lend += count
	bookDetail.Collection -= count
	models.GetGorm().Save(&bookDetail)

	if count > 0 {
		var borrow = models.Borrow{
			BookId: bookId,
			UserId: userId.(int),
			Count:  count,
		}
		models.GetGorm().Create(&borrow)
	} else {
		var borrow models.Borrow
		models.GetGorm().Where("book_id = ? AND user_id = ?", bookId, userId).First(&borrow)
		borrow.Count += count
		if borrow.Count == 0 {
			models.GetGorm().Delete(&borrow)
		} else {
			models.GetGorm().Save(&borrow)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"errmsg": "request OK",
	})
	return
}

func Manage(c *gin.Context) {
	userId, _ := c.Get("user_id")
	var user models.User
	models.GetGorm().Where("id = ?", userId).First(&user)
	if user.RoleId != 1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errmsg": "permission denied",
		})
		return
	}

	bookId, _ := strconv.Atoi(c.PostForm("book_id"))
	if bookId != 0 {
		var book models.BookDetail
		models.GetGorm().Where("id = ?", bookId).First(&book)
		if book.Lend != 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"errmsg": "book is lent",
			})
			return
		} else {
			models.GetGorm().Delete(&book)
			c.JSON(http.StatusOK, gin.H{
				"errmsg": "request OK",
			})
			return
		}
	}

	bookNum := c.PostForm("book_num")
	bookName := c.PostForm("book_name")
	author := c.PostForm("author")
	publisher := c.PostForm("publisher")
	category := c.PostForm("category")
	collection, _ := strconv.Atoi(c.PostForm("collection"))

	if bookNum == "" && bookName == "" && author == "" && publisher == "" && category == "" && collection == 0 {
		type userDetail struct {
			Name string `json:"name"`
			BookNum string `json:"book_num"`
			BookName string `json:"book_name"`
		}
		var userInfo []*userDetail
		models.GetGorm().Raw("SELECT users.name, book_details.book_num, book_details.book_name FROM users, book_details, borrows WHERE users.id = borrows.user_id AND book_details.id = borrows.book_id").Scan(&userInfo)
		c.JSON(http.StatusOK, gin.H{
			"errmsg": "request OK",
			"data": userInfo,
		})
		return
	}

	if bookNum == "" || bookName == "" || author == "" || publisher == "" || category == "" || collection == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"errmsg": "param missing",
		})
		return
	}

	var query models.BookDetail
	models.GetGorm().Where(
		"book_num = ?",
		bookNum,
	).First(&query)
	if query.Id != 0 && (bookName != query.BookName || author != query.Author || publisher != query.Publisher || category != query.Category) {
		c.JSON(http.StatusBadRequest, gin.H{
			"errmsg": "error, repeated book num",
		})
		return
	}

	if query.Id != 0 {
		query.Collection += collection
		models.GetGorm().Save(&query)
	} else if bookNum != "" && bookName != "" && author != "" && publisher != "" && category != "" && collection != 0 {
		var book = models.BookDetail{
			BookNum:    bookNum,
			BookName:   bookName,
			Author:     author,
			Publisher:  publisher,
			Category:   category,
			Collection: collection,
			Lend:       0,
		}
		models.GetGorm().Create(&book)
	}

	c.JSON(http.StatusOK, gin.H{
		"errmsg": "request OK",
	})
}
