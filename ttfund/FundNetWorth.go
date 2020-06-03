package ttfund

import (
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"regexp"
	"strconv"
	"strings"
)

//基金净值
type FundNetWorth struct {
	Date             string  //净值日期
	Unit             float32 //单位净值
	Accum            float32 //累计净值
	DailyGrowthRate  string  //日增长率
	PurchaseStatus   string  //申购状态
	RedemptionStatus string  //赎回状态
	Dividend         string  //分红送配
}

func parse(resp string) []FundNetWorth {
	var fnws []FundNetWorth

	trReg, _ := regexp.Compile("<tr>(.*?)</tr>")
	tdReg, _ := regexp.Compile("<td.*?>(.*?)</td>")

	trItems := trReg.FindAllString(resp, -1)
	for i := 0; i < len(trItems); i++ {
		trItem := trItems[i]
		tdItems := tdReg.FindAllString(trItem, -1)
		if nil != tdItems {
			if 7 == len(tdItems) {
				unit, _ := strconv.ParseFloat(parseTDItem(tdItems[1]), 32)
				accum, _ := strconv.ParseFloat(parseTDItem(tdItems[2]), 32)

				fundNW := FundNetWorth{
					Date:             parseTDItem(tdItems[0]),
					Unit:             float32(unit),
					Accum:            float32(accum),
					DailyGrowthRate:  parseTDItem(tdItems[3]),
					PurchaseStatus:   parseTDItem(tdItems[4]),
					RedemptionStatus: parseTDItem(tdItems[5]),
					Dividend:         parseTDItem(tdItems[6]),
				}
				fnws = append(fnws, fundNW)
			} else {
				fmt.Println(trItem)
			}
		}
	}

	return fnws
}

func parseTDItem(tdItem string) string {
	subItem := tdItem[0 : len(tdItem)-5]
	idx := strings.Index(subItem, ">")
	if idx < 0 {
		return ""
	}
	return subItem[idx+1:]
}

func fecthFundNetWorth(fundCode string) {
	hnwUrl := genHistoricalNetWorthUrl(fundCode, 1, 20, "20200526", "20200526")
	fmt.Println(hnwUrl)

	rawRes := common.HttpGet(hnwUrl)
	if rawRes == "" {
		return
	}
	fmt.Println(parse(rawRes))
}
