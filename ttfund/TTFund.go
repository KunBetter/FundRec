package ttfund

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"regexp"
	"strconv"
	"strings"
)

type TTFund struct {
}

func genHistoricalNetWorthUrl(fundCode int, pageIdx int, pageSize int, sDate string, eDate string) string {
	return fmt.Sprintf(hnwPrefix+"&code=%d&page=%d&per=%d&sdate=%s&edate=%s", fundCode, pageIdx, pageSize, sDate, eDate)
}

func (ttFund *TTFund) Test() {
	fetchFundValue("001186")
	fetchFundCompany()
	fetchFundList()
	fecthFundNetWorth(150270)
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

func fecthFundNetWorth(fundCode int) {
	hnwUrl := genHistoricalNetWorthUrl(fundCode, 1, 20, "20200526", "20200526")
	fmt.Println(hnwUrl)

	rawRes := common.HttpGet(hnwUrl)
	if rawRes == "" {
		return
	}
	fmt.Println(parse(rawRes))
}
