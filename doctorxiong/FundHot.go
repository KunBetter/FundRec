package doctorxiong

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
)

type FundHot struct {
	Code                  string `json:"code"`                  //基金代码
	Name                  string `json:"name"`                  //基金名称
	NetWorth              string `json:"netWorth"`              //当前基金单位净值
	DayGrowth             string `json:"dayGrowth"`             //日涨幅,单位为百分比
	LastWeekGrowth        string `json:"lastWeekGrowth"`        //最近一周涨幅,单位为百分比
	LastMonthGrowth       string `json:"lastMonthGrowth"`       //最近一个月涨幅,单位为百分比
	LastThreeMonthsGrowth string `json:"lastThreeMonthsGrowth"` //最近三个月涨幅,单位为百分比
	LastSixMonthsGrowth   string `json:"lastSixMonthsGrowth"`   //最近六个月涨幅,单位为百分比
	LastYearGrowth        string `json:"lastYearGrowth"`        //最近一年涨幅,单位为百分比
	ThisYearGrowth        string `json:"thisYearGrowth"`        //今年的涨幅,单位为百分比
}

type DXFundHotResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    []FundHot      `json:"data"`
	Meta    map[string]int `json:"meta"`
}

func fetchDXFundHot() {
	rawRes := common.HttpGet(dxFundHotUrl)
	if rawRes == "" {
		return
	}

	fv := &DXFundHotResponse{}
	err := json.Unmarshal([]byte(rawRes), &fv)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(fv)
}
