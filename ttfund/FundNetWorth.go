package ttfund

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
