package upload

import (
	"awesomeProject1/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct{
	id int `db:id`
	userid string `db:userid`
	img string `db:img`
}

func Upload(c *gin.Context)  {
	var use *User
	if use == nil{
		use = &User{}
	}
	type Json struct {
		userId string `form:"userId" json:"userId" binding:"required"`
	}
	j := Json{}
	//c.ShouldBindJSON(&j)
	j.userId = c.PostForm("userId")
	file, _ := c.FormFile("file")
	if _,error:=sql.DB.Exec("UPDATE  tb_user set img =? WHERE userid =?",file.Filename,j.userId); error!=nil{
		fmt.Errorf("query incur error")
	}else {
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
			return
		}
		c.JSON(200,gin.H{
			"msg":"上传成功",
			"img":file.Filename,
		})
		return
	}
	c.JSON(200,gin.H{
		"id":use.id,
		"userid":use.userid,
		"img":file.Filename,
	})
}