package entity

//基金持仓
type FundPosition struct {
	Title     string     `json:"title"`     //标题信息
	Date      string     `json:"date"`      //截至时间
	Stock     string     `json:"stock"`     //股票占比
	Bond      string     `json:"bond"`      //债券占比
	Cash      string     `json:"cash"`      //现金占比
	Total     string     `json:"total"`     //总净资产(亿元)
	StockList [][]string `json:"stockList"` //股票详情
}
