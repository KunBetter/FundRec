package core

import (
	"github.com/allegro/bigcache"
	"github.com/jinzhu/gorm"
	"time"
)

type FundRecCore struct {
	caches  map[int]*bigcache.BigCache
	mysqlDB *gorm.DB
}

func (frc *FundRecCore) Init() {
	frc.caches = make(map[int]*bigcache.BigCache)
	frc.addCache()
}

func (frc *FundRecCore) addCache() {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	frc.caches[0] = cache
	/*
		cache.Set("my-unique-key", []byte("value"))
		entry, _ := cache.Get("my-unique-key")
		fmt.Println(string(entry))
	*/
}

func (frc *FundRecCore) Run() {
	frc.FetchDXFundHot()
	frc.FetchDXFundDetail("000001")
	frc.FetchDXFund("202015")
	frc.FetchDXFundPosition("202015")
}
