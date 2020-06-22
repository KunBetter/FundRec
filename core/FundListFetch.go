package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
	"strings"
	"time"
)

type FundListFetch struct {
	RecCore *FundRecCore
}

func (fl *FundListFetch) Init() {
	fl.RecCore.mysqlDB.AutoMigrate(&entity.FundBaseInfo{})
}

func (fl *FundListFetch) Process() {
	curDay := time.Now().Format("2006-01-02")

	flag := GetFetchFlag(FundList)
	if nil != flag && curDay == flag.LatestDay {
		fmt.Println("FundList Fetched Today!")
		return
	}

	rawRes := common.HttpGet(common.TTFundsListUrl)
	if rawRes == "" {
		return
	}

	var fundsBuffer [][]string
	err := json.Unmarshal([]byte(rawRes[11:len(rawRes)-1]), &fundsBuffer)
	if err != nil {
		fmt.Println("some error")
	}

	var funds []entity.FundBaseInfo
	var fundName []string
	for i := 0; i < len(fundsBuffer); i++ {
		fundBuffer := fundsBuffer[i]
		if len(fundBuffer) == 5 {
			fund := entity.FundBaseInfo{
				Code:        fundBuffer[0],
				Name:        fundBuffer[2],
				Type:        fundBuffer[3],
				SingleSpell: fundBuffer[1],
				AllSpell:    fundBuffer[4],
			}
			funds = append(funds, fund)
			fundName = append(fundName, fund.Name)
		} else {
			fmt.Println(fundBuffer)
		}
	}

	flushMD5 := Hash(strings.Join(fundName, ""))
	if flag != nil && flag.MD5 == flushMD5 {
		fmt.Println("FundList Data Not Changed!")
		return
	}

	if nil == flag {
		flag = &FetchFlag{}
	}
	flag.MD5 = flushMD5
	flag.LatestDay = curDay

	fl.RecCore.DBRef().Delete(entity.FundBaseInfo{})

	for i := 0; i < len(funds); i++ {
		fl.RecCore.DBRef().Create(&funds[i])
	}

	WriteFlag(FundList, flag)
}
