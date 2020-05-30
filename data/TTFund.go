package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	/**
	 * 获取基金历史净值地址
	 * Historical Net Worth
	 * 参数：&code=150270&page=1&per=20&sdate=20200526&edate=20200526"
	 */
	hnwPrefix = "http://fund.eastmoney.com/f10/F10DataApi.aspx?type=lsjz"
	/**
	 * 获取基金列表
	 */
	fundsUrl = "http://fund.eastmoney.com/js/fundcode_search.js"
)

func genHistoricalNetWorthUrl(fundCode int, pageIdx int, pageSize int, sDate string, eDate string) string {
	return fmt.Sprintf(hnwPrefix+"&code=%d&page=%d&per=%d&sdate=%s&edate=%s", fundCode, pageIdx, pageSize, sDate, eDate)
}

func main() {
	fetchFundList()
	fecthFundNetWorth(150270)
}

func fetchFundList() {
	resp, err := http.Get(fundsUrl)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	if 200 == resp.StatusCode {
		buf := bytes.NewBuffer(make([]byte, 0, 512))
		length, _ := buf.ReadFrom(resp.Body)

		if len(buf.Bytes()) == int(length) {
			rawRes := string(buf.Bytes())

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
	}
}

type Fund struct {
	Code        string //基金代码
	SingleSpell string //基金名称单拼
	Name        string //基金名称中文
	Type        string //基金类型
	AllSpell    string //基金名称全拼
}

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

func fecthFundNetWorth(fundCode int) {
	hnwUrl := genHistoricalNetWorthUrl(fundCode, 1, 20, "20200526", "20200526")
	fmt.Println(hnwUrl)
	resp, err := http.Get(hnwUrl)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	if 200 == resp.StatusCode {
		buf := bytes.NewBuffer(make([]byte, 0, 512))
		length, _ := buf.ReadFrom(resp.Body)

		if len(buf.Bytes()) == int(length) {
			rawRes := string(buf.Bytes())
			fmt.Println(parse(rawRes))
		}
	}
}
