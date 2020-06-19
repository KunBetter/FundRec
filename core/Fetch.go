package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
	"io/ioutil"
	"strings"
)

const (
	/*
	 * md5
	 * latest time
	 */
	FundCompanyFlag = "FundCompany.md5"
)

type Fetch interface {
	CreateTable()
	Process()
	WriteFlag()
	Store(data []interface{})
}

type Flag struct {
	MD5        string
	LatestTime string
}

func GetFetchFlag() *Flag {
	buf, err := ioutil.ReadFile(FundCompanyFlag)
	if err != nil {
		fmt.Print(err)
	}
	lines := strings.Split(string(buf), "\n")
	if 2 == len(lines) {
		flag := &Flag{
			MD5:        lines[0],
			LatestTime: lines[1],
		}

		return flag
	}

	return nil
}

func WriteFlag() {
}

type FundCompanyFetch struct {
	engine *FundRecCore
}

func (fc *FundCompanyFetch) CreateTable() {
	fc.engine.mysqlDB.AutoMigrate(&entity.FundCompany{})
}

func (fc *FundCompanyFetch) Process() {
	rawRes := common.HttpGet(common.TTFundCompanyUrl)
	if rawRes == "" {
		return
	}

	var fundCompanyBuffer [][]string
	err := json.Unmarshal([]byte(rawRes[14:len(rawRes)-1]), &fundCompanyBuffer)
	if err != nil {
		fmt.Println("some error")
	}

	var fundCompanies []entity.FundCompany
	for i := 0; i < len(fundCompanyBuffer); i++ {
		fcBuffer := fundCompanyBuffer[i]
		if len(fcBuffer) == 2 {
			fundCompany := entity.FundCompany{
				Code: fcBuffer[0],
				Name: fcBuffer[1],
			}
			fundCompanies = append(fundCompanies, fundCompany)
		} else {
			fmt.Println(fcBuffer)
		}
	}

	for i := 0; i < len(fundCompanies); i++ {
		fc.engine.DBRef().Create(&fundCompanies[i])
	}
}

func (fc *FundCompanyFetch) Store(data []interface{}) {
}

func (fc *FundCompanyFetch) WriteFlag() {
}
