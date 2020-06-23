package core

import (
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/config"
	"github.com/KunBetter/FundRec/strategy"
	"github.com/allegro/bigcache"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type FundRecCore struct {
	conf    *config.Config
	mysqlDB *gorm.DB
	caches  map[string]*bigcache.BigCache

	fcFetch *FundCompanyFetch
	flFetch *FundListFetch
	fvFetch *FundValueFetch

	rs *strategy.RecStrategy
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

	frc.fvFetch = &FundValueFetch{
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

func (frc *FundRecCore) FundDataFetch() {
	frc.fcFetch.Start()
	frc.flFetch.Start()
	frc.fvFetch.Start()
}

func (frc *FundRecCore) Rec() {
}

func (frc *FundRecCore) Router() {
	router := gin.Default()

	router.GET("/rec/funds", frc.GetRecFunds)
	router.Run(":8080")
}

func (frc *FundRecCore) GetRecFunds(c *gin.Context) {
	frc.Rec()
}
