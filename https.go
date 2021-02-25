package main

import (
	"awesomeProject1/Test"
	"awesomeProject1/model"
	"awesomeProject1/upload"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)
var CRX *Test.Book
func main ()  {
	if CRX == nil {
		CRX = Test.Curse()
	}

	f,_ :=os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//测试api是否正常访问
	sql.Db()
	Upload := router.Group("/v3")
	{
		Upload.POST("/upload",upload.Upload)
	}
	router.GET("/ping", CRX.Create)
	router.Run(":8085")
	CRX.Set()
	countryCapitalMap := make(map[string]string)
	countryCapitalMap["测试"] = "测试"
	fmt.Println(countryCapitalMap["测试"])
}


