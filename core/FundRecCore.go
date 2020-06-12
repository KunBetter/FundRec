package core

import (
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/allegro/bigcache"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type FundRecCore struct {
	caches  map[string]*bigcache.BigCache
	mysqlDB *gorm.DB
}

func (frc *FundRecCore) DBRef() *gorm.DB {
	return frc.mysqlDB
}

func (frc *FundRecCore) Init() bool {
	frc.caches = make(map[string]*bigcache.BigCache)
	frc.addCache()

	var err error
	frc.mysqlDB, err = gorm.Open("mysql", "root:root1234@tcp(127.0.0.1:3306)/funds?charset=utf8")
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return false
	} else {
		fmt.Println("connect database success")
		frc.mysqlDB.SingularTable(true)
	}
	//defer frc.mysqlDB.Close()

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
	frc.FetchRankedFunds()
	//frc.FetchDXFundDetail("001186")
}
