package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
)

func (frc *FundRecCore) FetchFundCompany() {
	//Create Table
	frc.mysqlDB.AutoMigrate(&entity.FundCompany{})

	rawRes := common.HttpGet(common.TTFundCompanyUrl)
	if rawRes == "" {
		return
	}

	var fundCompanyBuffer [][]string
	err := json.Unmarshal([]byte(rawRes[14:len(rawRes)-1]), &fundCompanyBuffer)
	if err != nil {
		fmt.Println("some error")
	}

	for i := 0; i < len(fundCompanyBuffer); i++ {
		fcBuffer := fundCompanyBuffer[i]
		if len(fcBuffer) == 2 {
			fundCompany := entity.FundCompany{
				Code: fcBuffer[0],
				Name: fcBuffer[1],
			}
			//Insert to DB
			frc.DBRef().Create(&fundCompany)
		} else {
			fmt.Println(fcBuffer)
		}
	}
}
