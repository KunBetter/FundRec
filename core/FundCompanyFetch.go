package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
	"github.com/KunBetter/FundRec/schedule"
	"strings"
	"time"
)

type FundCompanyFetch struct {
	RecCore   *FundRecCore
	Scheduler *schedule.Scheduler
}

func (fc *FundCompanyFetch) Init() {
	fc.RecCore.DBRef().AutoMigrate(&entity.FundCompany{})
}

func (fc *FundCompanyFetch) Start() {
	//engine start
	fc.Init()
	fc.Process()

	fc.Scheduler = schedule.NewScheduler(BaseFetchSpecTime, fc)
	fc.Scheduler.Start()
}

func (fc *FundCompanyFetch) Run() {
	fc.Process()
}

func (fc *FundCompanyFetch) Process() {
	curDay := time.Now().Format("2006-01-02")

	flag := GetFetchFlag(FundCompany)
	if nil != flag && curDay == flag.LatestDay {
		fmt.Println("FundCompany Fetched Today!")
		return
	}

	rawRes := common.HttpGet(common.TTFundCompanyUrl)
	if rawRes == "" {
		return
	}

	var fundCompanyBuffer [][]string
	err := json.Unmarshal([]byte(rawRes[14:len(rawRes)-1]), &fundCompanyBuffer)
	if err != nil {
		fmt.Println("some error")
	}

	if len(fundCompanyBuffer) <= 0 {
		return
	}

	var fundCompanies []entity.FundCompany
	var companyName []string
	for i := 0; i < len(fundCompanyBuffer); i++ {
		fcBuffer := fundCompanyBuffer[i]
		if len(fcBuffer) == 2 {
			fundCompany := entity.FundCompany{
				Code: fcBuffer[0],
				Name: fcBuffer[1],
			}
			fundCompanies = append(fundCompanies, fundCompany)
			companyName = append(companyName, fundCompany.Name)
		} else {
			fmt.Println(fcBuffer)
		}
	}

	flushMD5 := Hash(strings.Join(companyName, ""))
	if flag != nil && flag.MD5 == flushMD5 {
		fmt.Println("FundCompany Data Not Changed!")
		return
	}

	if nil == flag {
		flag = &FetchFlag{}
	}
	flag.MD5 = flushMD5
	flag.LatestDay = curDay

	fc.RecCore.DBRef().Delete(entity.FundCompany{})

	for i := 0; i < len(fundCompanies); i++ {
		fc.RecCore.DBRef().Create(&fundCompanies[i])
	}

	WriteFlag(FundCompany, flag)
}
