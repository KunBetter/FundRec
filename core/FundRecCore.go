package core

import (
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/config"
	"github.com/allegro/bigcache"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
