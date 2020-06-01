package doctorxiong

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
)

/**
 * https://www.doctorxiong.club
 */

const (
	dxFundUrl = "https://api.doctorxiong.club/v1/fund?code=%s"
)

type DXFund struct {
}

func (dxf *DXFund) Test() {
	fetchDXFund("202015")
}

func fetchDXFund(code string) {
	rawRes := common.HttpGet(fmt.Sprintf(dxFundUrl, code))
	if rawRes == "" {
		return
	}

	// TODO have some Problem
	fv := &DXFResponse{}
	err := json.Unmarshal([]byte(rawRes), &fv)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(fv)
}

type DXFResponse struct {
	code    int
	message string
	data    []Fund
	meta    string
}

type Fund struct {
	Code                    string  `json:"code"`                    //基金代码
	Name                    string  `json:"name"`                    //基金名称
	Type                    string  `json:"type"`                    //基金类型
	NetWorth                float32 `json:"netWorth"`                //当前基金单位净值
	ExpectWorth             float32 `json:"expectWorth"`             //当前基金单位净值估算
	TotalWorth              float32 `json:"totalWorth"`              //当前基金累计净值
	ExpectGrowth            string  `json:"expectGrowth"`            //当前基金单位净值估算日涨幅,单位为百分比
	DayGrowth               string  `json:"dayGrowth"`               //单位净值日涨幅,单位为百分比
	LastWeekGrowth          string  `json:"lastWeekGrowth"`          //单位净值周涨幅,单位为百分比
	LastMonthGrowth         string  `json:"lastMonthGrowth"`         //单位净值月涨幅,单位为百分比
	LastThreeMonthsGrowth   string  `json:"lastThreeMonthsGrowth"`   //单位净值三月涨幅,单位为百分比
	LastSixMonthsGrowth     string  `json:"lastSixMonthsGrowth"`     //单位净值六月涨幅,单位为百分比
	LastYearGrowth          string  `json:"lastYearGrowth"`          //单位净值年涨幅,单位为百分比
	BuyMin                  float32 `json:"buyMin"`                  //起购额度
	BuySourceRate           float32 `json:"buySourceRate"`           //原始买入费率,单位为百分比
	BuyRate                 float32 `json:"buyRate"`                 //当前买入费率,单位为百分比
	Manager                 string  `json:"manager"`                 //基金经理
	FundScale               string  `json:"fund_Scale"`              //基金规模及日期,日期为最后一次规模变动的日期
	WorthDate               string  `json:"worthDate"`               //净值更新日期,日期格式为yy-MM-dd HH:mm.2019-06-27 15:00代表当天下午3点
	ExpectWorthDate         string  `json:"expectWorthDate"`         //净值估算更新日期,,日期格式为yy-MM-dd HH:mm.2019-06-27 15:00代表当天下午3点
	MillionCopiesIncome     float32 `json:"millionCopiesIncome"`     //每万分收益(货币基金)
	MillionCopiesIncomeDate string  `json:"millionCopiesIncomeDate"` //七日年化收益更新日期(货币基金)
	SevenDaysYearIncome     float32 `json:"sevenDaysYearIncome"`     //七日年化收益(货币基金)
}
