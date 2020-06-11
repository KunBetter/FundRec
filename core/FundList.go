package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
)

func (frc *FundRecCore) FetchFundList() {
	rawRes := common.HttpGet(common.TTFundsListUrl)
	if rawRes == "" {
		return
	}

	var fundsBuffer [][]string
	err := json.Unmarshal([]byte(rawRes[11:len(rawRes)-1]), &fundsBuffer)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(len(fundsBuffer))

	var funds []entity.FundBaseInfo
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
		} else {
			fmt.Println(fundBuffer)
		}
	}
	fmt.Println(funds)
}
