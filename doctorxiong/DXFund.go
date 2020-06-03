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
	/**
	 * GetV1Fund - 获取基金基础信息
	 */
	dxFundUrl = "https://api.doctorxiong.club/v1/fund?code=%s"

	/**
	 * GetV1FundPosition - 获取基金持仓详情
	 */
	dxFundPosUrl = "https://api.doctorxiong.club/v1/fund/position?code=%s"

	/**
	 * PostV1FundRank - 获取基金排行
	 */
	dxFundRankUrl = "https://api.doctorxiong.club/v1/fund/rank"

	/**
	 * GetV1FundDetail - 获取基金详情
	 */
	dxFundDetailUrl = "https://api.doctorxiong.club/v1/fund/detail?code=%s"

	/**
	 * GetV1FundHot - 获取热门基金
	 */
	dxFundHotUrl = "https://api.doctorxiong.club/v1/fund/hot"
)

type DXFund struct {
}

func (dxf *DXFund) Test() {
	fetchDXFundDetail("000001")
	fetchDXFund("202015")
	fetchDXFundPosition("202015")
}

func fetchDXFundDetail(code string) {
	rawRes := common.HttpGet(fmt.Sprintf(dxFundDetailUrl, code))
	if rawRes == "" {
		return
	}

	fv := &DXFundDetailResponse{}
	err := json.Unmarshal([]byte(rawRes), &fv)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(fv)
}

type DXFundDetailResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    FundDetail `json:"data"`
	Meta    string     `json:"meta"`
}

type FundDetail struct {
	Code                    string          `json:"code"`                    //基金代码
	Name                    string          `json:"name"`                    //基金名称
	Type                    string          `json:"type"`                    //基金类型
	NetWorth                float32         `json:"netWorth"`                //当前基金单位净值
	ExpectWorth             float32         `json:"expectWorth"`             //当前基金单位净值估算
	TotalWorth              float32         `json:"totalWorth"`              //当前基金累计净值
	ExpectGrowth            string          `json:"expectGrowth"`            //当前基金单位净值估算日涨幅,单位为百分比
	DayGrowth               string          `json:"dayGrowth"`               //单位净值日涨幅,单位为百分比
	LastWeekGrowth          string          `json:"lastWeekGrowth"`          //单位净值周涨幅,单位为百分比
	LastMonthGrowth         string          `json:"lastMonthGrowth"`         //单位净值月涨幅,单位为百分比
	LastThreeMonthsGrowth   string          `json:"lastThreeMonthsGrowth"`   //单位净值三月涨幅,单位为百分比
	LastSixMonthsGrowth     string          `json:"lastSixMonthsGrowth"`     //单位净值六月涨幅,单位为百分比
	LastYearGrowth          string          `json:"lastYearGrowth"`          //单位净值年涨幅,单位为百分比
	BuyMin                  float32         `json:"buyMin"`                  //起购额度
	BuySourceRate           float32         `json:"buySourceRate"`           //原始买入费率,单位为百分比
	BuyRate                 float32         `json:"buyRate"`                 //当前买入费率,单位为百分比
	Manager                 string          `json:"manager"`                 //基金经理
	FundScale               string          `json:"fundScale"`               //基金规模及日期,日期为最后一次规模变动的日期
	WorthDate               string          `json:"worthDate"`               //净值更新日期,日期格式为yy-MM-dd HH:mm.2019-06-27 15:00代表当天下午3点
	ExpectWorthDate         string          `json:"expectWorthDate"`         //净值估算更新日期,,日期格式为yy-MM-dd HH:mm.2019-06-27 15:00代表当天下午3点
	NetWorthData            [][]interface{} `json:"netWorthData"`            //历史净值信息["2001-12-18" , 1 , 0 , ""]依次表示:日期; 单位净值; 净值涨幅; 每份分红.
	MillionCopiesIncome     float32         `json:"millionCopiesIncome"`     //每万分收益(货币基金)
	MillionCopiesIncomeData [][]interface{} `json:"millionCopiesIncomeData"` //历史万每分收益信息(货币基金)["2016-09-23",0.4773]依次表示:日期; 每万分收益.
	MillionCopiesIncomeDate string          `json:"millionCopiesIncomeDate"` //七日年化收益更新日期(货币基金)
	SevenDaysYearIncome     float32         `json:"sevenDaysYearIncome"`     //七日年化收益(货币基金)
	SevenDaysYearIncomeData [][]interface{} `json:"sevenDaysYearIncomeData"` //历史七日年华收益信息(货币基金)["2016-09-23",2.131]依次表示:日期; 七日年化收益.
}

func fetchDXFundPosition(code string) {
	rawRes := common.HttpGet(fmt.Sprintf(dxFundPosUrl, code))
	if rawRes == "" {
		return
	}

	fv := &DXFPosResponse{}
	err := json.Unmarshal([]byte(rawRes), &fv)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(fv)
}

func fetchDXFund(code string) {
	rawRes := common.HttpGet(fmt.Sprintf(dxFundUrl, code))
	if rawRes == "" {
		return
	}

	fv := &DXFResponse{}
	err := json.Unmarshal([]byte(rawRes), &fv)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(fv)
}

type DXFResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Fund `json:"data"`
	Meta    string `json:"meta"`
}

type DXFPosResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    FundPos `json:"data"`
	Meta    string  `json:"meta"`
}

type FundPos struct {
	Title     string     `json:"title"`     //标题信息
	Date      string     `json:"date"`      //截至时间
	Stock     string     `json:"stock"`     //股票占比
	Bond      string     `json:"bond"`      //债券占比
	Cash      string     `json:"cash"`      //现金占比
	Total     string     `json:"total"`     //总净资产(亿元)
	StockList [][]string `json:"stockList"` //股票详情
}

type Stock struct {
	Code  string //股票代码
	Name  string //名称
	Ratio string //占比
	Num   string //持有股数(万股)
	Cash  string //持有金额(万元)
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
	FundScale               string  `json:"fundScale"`               //基金规模及日期,日期为最后一次规模变动的日期
	WorthDate               string  `json:"worthDate"`               //净值更新日期,日期格式为yy-MM-dd HH:mm.2019-06-27 15:00代表当天下午3点
	ExpectWorthDate         string  `json:"expectWorthDate"`         //净值估算更新日期,,日期格式为yy-MM-dd HH:mm.2019-06-27 15:00代表当天下午3点
	MillionCopiesIncome     float32 `json:"millionCopiesIncome"`     //每万分收益(货币基金)
	MillionCopiesIncomeDate string  `json:"millionCopiesIncomeDate"` //七日年化收益更新日期(货币基金)
	SevenDaysYearIncome     float32 `json:"sevenDaysYearIncome"`     //七日年化收益(货币基金)
}
