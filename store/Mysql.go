package store

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type FundStore struct {
	MysqlDB *gorm.DB
}

type User struct {
	Id   int    `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Age  int    `gorm:"size:11;DEFAULT NULL" json:"age"`
	Name string `gorm:"size:255;DEFAULT NULL" json:"name"`
}

func (fs *FundStore) Store() {
	var err error
	fs.MysqlDB, err = gorm.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/funds?charset=utf8")
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return
	} else {
		fmt.Println("connect database success")
		fs.MysqlDB.SingularTable(true)
		fs.MysqlDB.AutoMigrate(&User{}) //自动建表
		fmt.Println("create table success")
	}
	defer fs.MysqlDB.Close()

	fs.Router()
}

func (fs *FundStore) Router() {
	router := gin.Default()
	//路径映射
	router.GET("/user", fs.InitPage)
	router.POST("/user/create", fs.CreateUser)
	router.GET("/user/list", fs.ListUser)
	router.PUT("/user/update", fs.UpdateUser)
	router.GET("/user/find", fs.GetUser)
	router.DELETE("/user/:id", fs.DeleteUser)
	router.Run(":8080")
}

//每个路由都对应一个具体的函数操作,从而实现了对user的增,删,改,查操作
func (fs *FundStore) InitPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
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
