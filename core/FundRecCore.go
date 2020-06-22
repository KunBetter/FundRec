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

	fcFetch *FundCompanyFetch
	flFetch *FundListFetch
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

	frc.fcFetch = &FundCompanyFetch{
		RecCore: frc,
	}

	frc.flFetch = &FundListFetch{
		RecCore: frc,
	}

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

func (frc *FundRecCore) Router() {
	router := gin.Default()

	router.POST("/rec/create", frc.PostFunc)
	router.PUT("/rec/update", frc.PutFunc)
	router.GET("/rec/find", frc.GetFunc)

	router.Run(":8080")
}

func (frc *FundRecCore) PostFunc(c *gin.Context) {
	var user interface{}
	c.BindJSON(&user) //使用bindJSON填充对象
	//biz code
	c.JSON(http.StatusOK, &user) //返回页面
}

func (frc *FundRecCore) PutFunc(c *gin.Context) {
	var user interface{}
	c.PostForm("id")
	c.BindJSON(&user)
	//biz code
	c.JSON(http.StatusOK, &user)
}

func (frc *FundRecCore) GetFunc(c *gin.Context) {
	line := c.Query("line")
	//biz code
	c.JSON(http.StatusOK, &line)
}

func (frc *FundRecCore) FundDataFetch() {
	frc.fcFetch.Init()
	frc.fcFetch.Process()

	frc.flFetch.Init()
	frc.flFetch.Process()

	/*
		go frc.FetchHotFunds()
		go frc.FetchRankedFunds()

		go frc.FetchFundNetWorth("150270")
		go frc.FetchFundValue("001186")
		go frc.FetchFundPosition("001186")
		go frc.FetchFund("202015")
		go frc.FetchDXFundDetail("003171")
	*/
}
