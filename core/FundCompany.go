package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
)

func (frc *FundRecCore) FetchFundCompany() {
	rawRes := common.HttpGet(common.TTFundCompanyUrl)
	if rawRes == "" {
		return
	}

	var fundCompanyBuffer [][]string
	err := json.Unmarshal([]byte(rawRes[14:len(rawRes)-1]), &fundCompanyBuffer)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(len(fundCompanyBuffer))

	var fundCompanys []entity.FundCompany
	for i := 0; i < len(fundCompanyBuffer); i++ {
		fcBuffer := fundCompanyBuffer[i]
		if len(fcBuffer) == 2 {
			fundCompany := entity.FundCompany{
				Code: fcBuffer[0],
				Name: fcBuffer[1],
			}
			fundCompanys = append(fundCompanys, fundCompany)
		} else {
			fmt.Println(fcBuffer)
		}
	}
	fmt.Println(fundCompanys)
}
