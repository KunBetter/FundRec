package ttfund

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
)

type Fund struct {
	Code        string //基金代码
	SingleSpell string //基金名称单拼
	Name        string //基金名称中文
	Type        string //基金类型
	AllSpell    string //基金名称全拼
}

func fetchFundList() {
	rawRes := common.HttpGet(fundsUrl)
	if rawRes == "" {
		return
	}

	var fundsBuffer [][]string
	err := json.Unmarshal([]byte(rawRes[11:len(rawRes)-1]), &fundsBuffer)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(len(fundsBuffer))

	var funds []Fund
	for i := 0; i < len(fundsBuffer); i++ {
		fundBuffer := fundsBuffer[i]
		if len(fundBuffer) == 5 {
			fund := Fund{
				Code:        fundBuffer[0],
				SingleSpell: fundBuffer[1],
				Name:        fundBuffer[2],
				Type:        fundBuffer[3],
				AllSpell:    fundBuffer[4],
			}
			funds = append(funds, fund)
		} else {
			fmt.Println(fundBuffer)
		}
	}
	fmt.Println(funds)
}
