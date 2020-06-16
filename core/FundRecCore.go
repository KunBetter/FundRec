package core

import (
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/config"
	"github.com/allegro/bigcache"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"time"
)

type FundRecCore struct {
	conf    *config.Config
	mysqlDB *gorm.DB
	caches  map[string]*bigcache.BigCache
}

func (frc *FundRecCore) DBRef() *gorm.DB {
	return frc.mysqlDB
}

func (frc *FundRecCore) Init() bool {
	frc.conf = config.LoadConfig()

	DBUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		frc.conf.Mysql.Username, frc.conf.Mysql.Password, frc.conf.Mysql.Host, frc.conf.Mysql.Port, frc.conf.Mysql.DBName)

	var err error
	frc.mysqlDB, err = gorm.Open("mysql", DBUri)
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return false
	} else {
		fmt.Println("connect database success")
		frc.mysqlDB.SingularTable(true)
	}
	//defer frc.mysqlDB.Close()

	frc.caches = make(map[string]*bigcache.BigCache)
	frc.addCache()

	return true
}

func (frc *FundRecCore) addCache() {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	frc.caches[common.FundCompanyCache] = cache
	/*
		cache.Set("my-unique-key", []byte("value"))
		entry, _ := cache.Get("my-unique-key")
		fmt.Println(string(entry))
	*/
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

func (frc *FundRecCore) Run() {
	//frc.FetchFundCompany()
	//frc.FetchFundList()
	//frc.FetchFundNetWorth("150270")
	//frc.FetchFundValue("001186")
	//frc.FetchFundPosition("001186")
	//frc.FetchFund("202015")
	//frc.FetchHotFunds()
	//frc.FetchRankedFunds()
	frc.FetchDXFundDetail("003171")
}
