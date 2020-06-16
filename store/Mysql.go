package store

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type MysqlDB struct {
	db *gorm.DB
}


func (ptr *MysqlDB) Open(dbUri string) {
	var err error
	ptr.db, err = gorm.Open("mysql", dbUri)
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return
	} else {
		fmt.Println("connect database success")
		ptr.db.SingularTable(true)
	}
	defer ptr.db.Close()
}

func (ptr *MysqlDB)Insert(){

}


func (fs *FundStore) CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)            //使用bindJSON填充对象
	fs.MysqlDB.Create(&user)     //创建对象
	c.JSON(http.StatusOK, &user) //返回页面
}

func (fs *FundStore) UpdateUser(c *gin.Context) {
	var user User
	id := c.PostForm("id")                   //post方法取相应字段
	err := fs.MysqlDB.First(&user, id).Error //数据库查找主键=ID的第一行
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err.Error())
	} else {
		c.BindJSON(&user)
		fs.MysqlDB.Save(&user) //提交更改
		c.JSON(http.StatusOK, &user)
	}
}

func (fs *FundStore) ListUser(c *gin.Context) {
	var user []User
	line := c.Query("line")
	fs.MysqlDB.Limit(line).Find(&user) //限制查找前line行
	c.JSON(http.StatusOK, &user)
}

func (fs *FundStore) GetUser(c *gin.Context) {
	id := c.Query("id")
	var user User
	err := fs.MysqlDB.First(&user, id).Error
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err.Error())
	} else {
		c.JSON(http.StatusOK, &user)
	}
}

func (fs *FundStore) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	fs.MysqlDB.Where("id = ?", id).Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"data": "this has been deleted!",
	})
}
