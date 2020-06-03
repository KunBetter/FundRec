package ttfund

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
)

type FundValue struct {
	Code                string `json:"fundcode"` //基金代码
	Name                string `json:"name"`     //基金名称
	NetWorthDate        string `json:"jzrq"`     //净值日期
	NetValueCurDay      string `json:"dwjz"`     //当日净值
	EstimatedNetWorth   string `json:"gsz"`      //估算净值
	EstimatedPercentage string `json:"gszzl"`    //估算涨跌百分比
	EstimateTime        string `json:"gztime"`   //估值时间
}

func fetchFundValue(code string) {
	rawRes := common.HttpGet(fmt.Sprintf(fundValueUrl, code))
	if rawRes == "" {
		return
	}

	fv := &FundValue{}
	err := json.Unmarshal([]byte(rawRes[8:len(rawRes)-2]), &fv)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(fv)
}
