package doctorxiong

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
)

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

type DXFPosResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    FundPos `json:"data"`
	Meta    string  `json:"meta"`
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
