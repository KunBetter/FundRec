package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const (
	/**
	 * 获取基金历史净值地址
	 * Historical Net Worth
	 * 参数：&code=150270&page=1&per=20&sdate=20200526&edate=20200526"
	 */
	hnwPrefix = "http://fund.eastmoney.com/f10/F10DataApi.aspx?type=lsjz"
)

func genHistoricalNetWorthUrl(fundCode int, pageIdx int, pageSize int, sDate string, eDate string) string {
	return fmt.Sprintf(hnwPrefix+"&code=%d&page=%d&per=%d&sdate=%s&edate=%s", fundCode, pageIdx, pageSize, sDate, eDate)
}

func main() {
	fund()
}

/*
 * 1、获取基金列表
 * 2、单个基金净值
 * 3、基金排名
 */

//基金净值
type FundNetWorth struct {
	Date             string  //净值日期
	Unit             float32 //单位净值
	Accum            float32 //累计净值
	DailyGrowthRate  float32 //日增长率
	PurchaseStatus   string  //申购状态
	RedemptionStatus string  //赎回状态
	Dividend         string  //分红送配
}

func parse(resp string) []*FundNetWorth {
	var fnws []*FundNetWorth

	trReg, _ := regexp.Compile("<tr>(.*?)</tr>")
	tdReg, _ := regexp.Compile("<td.*?>(.*?)</td>")

	trItems := trReg.FindAllString(resp, -1)
	for i := 0; i < len(trItems); i++ {
		trItem := trItems[i]
		tdItems := tdReg.FindAllString(trItem, -1)
		if nil != tdItems {
			if 7 == len(tdItems) {
				fundNW := &FundNetWorth{
					Date:             "",
					Unit:             1.0,
					Accum:            1.0,
					DailyGrowthRate:  1.0,
					PurchaseStatus:   "",
					RedemptionStatus: "",
					Dividend:         "",
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

func fund() {
	hnwUrl := genHistoricalNetWorthUrl(150270, 1, 20, "20200526", "20200526")
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
