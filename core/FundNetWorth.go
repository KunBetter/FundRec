package core

import (
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
	"regexp"
	"strconv"
	"strings"
)

func (frc *FundRecCore) FecthFundNetWorth(fundCode string) { //150270
	hnwUrl := genHistoricalNetWorthUrl(fundCode, 1, 20, "20200526", "20200526")
	fmt.Println(hnwUrl)

	rawRes := common.HttpGet(hnwUrl)
	if rawRes == "" {
		return
	}
	fmt.Println(parse(rawRes))
}

func parse(resp string) []entity.FundNetWorth {
	var fnws []entity.FundNetWorth

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

				fundNW := entity.FundNetWorth{
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

func genHistoricalNetWorthUrl(fundCode string, pageIdx int, pageSize int, sDate string, eDate string) string {
	return fmt.Sprintf(common.HNWPrefix+"&code=%s&page=%d&per=%d&sdate=%s&edate=%s", fundCode, pageIdx, pageSize, sDate, eDate)
}
