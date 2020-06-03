package ttfund

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
)

type FundCompany struct {
	Code string //基金公司代码
	Name string //基金公司名称中文
}

func fetchFundCompany() {
	rawRes := common.HttpGet(fundCompanyUrl)
	if rawRes == "" {
		return
	}

	var fundCompanyBuffer [][]string
	err := json.Unmarshal([]byte(rawRes[14:len(rawRes)-1]), &fundCompanyBuffer)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(len(fundCompanyBuffer))

	var fundCompanys []FundCompany
	for i := 0; i < len(fundCompanyBuffer); i++ {
		fcBuffer := fundCompanyBuffer[i]
		if len(fcBuffer) == 2 {
			fundCompany := FundCompany{
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
