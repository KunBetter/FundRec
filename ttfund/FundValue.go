package ttfund

type FundValue struct {
	Code                string `json:"fundcode"` //基金代码
	Name                string `json:"name"`     //基金名称
	NetWorthDate        string `json:"jzrq"`     //净值日期
	NetValueCurDay      string `json:"dwjz"`     //当日净值
	EstimatedNetWorth   string `json:"gsz"`      //估算净值
	EstimatedPercentage string `json:"gszzl"`    //估算涨跌百分比
	EstimateTime        string `json:"gztime"`   //估值时间
}
