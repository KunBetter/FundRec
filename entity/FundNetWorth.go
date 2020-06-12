package entity

//基金净值
type FundNetWorth struct {
	Id               int     `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Code             string  `gorm:"size:127;DEFAULT NULL" json:"code"`              //基金代码
	Date             string  `gorm:"size:127;DEFAULT NULL" json:"date"`              //净值日期
	Unit             float32 `gorm:"type:decimal(20,4);DEFAULT 0.0" json:"unit"`     //单位净值
	Accum            float32 `gorm:"type:decimal(20,4);DEFAULT 0.0" json:"accum"`    //累计净值
	DailyGrowthRate  string  `gorm:"size:127;DEFAULT NULL" json:"daily_growth_rate"` //日增长率；净值涨幅
	PurchaseStatus   string  `gorm:"size:127;DEFAULT NULL" json:"purchase_status"`   //申购状态
	RedemptionStatus string  `gorm:"size:127;DEFAULT NULL" json:"redemption_status"` //赎回状态
	Dividend         string  `gorm:"size:127;DEFAULT NULL" json:"dividend"`          //分红送配；每份分红
}
