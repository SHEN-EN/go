package Test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	Name int
	Title int
}
type change interface { // 定义接口
	set()
}

func (book Book) Set()  { // 实现接口方法
	fmt.Printf("123")
}
func Curse() *Book {
	c := Book{Name: 1,Title: 2}
	return &c
}
func (book *Book) Create(c *gin.Context)  {
	book.Name = 0
	book.Title = 1
	c.JSON(http.StatusOK,gin.H{
		"msg":  book,
	})
}

func (book Book) StayTheSame() {
	book.Name = 5
	book.Title = 6
	fmt.Print(book)
}

