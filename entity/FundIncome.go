package entity

//货币基金
//1、历史万每分收益信息
//2、历史七日年华收益信息
type FundIncome struct {
	Id   int    `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Code string `gorm:"size:127;DEFAULT NULL" json:"code"`
	//MillionCopiesIncome,每万分收益,1
	//SevenDaysYearIncome,七日年华收益,2
	Type   int     `gorm:"type:int(11);;DEFAULT 0" json:"type"`
	Date   string  `gorm:"size:127;DEFAULT NULL" json:"date"`
	Income float32 `gorm:"type:decimal(20,4);DEFAULT 0.0" json:"income"`
}
