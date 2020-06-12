package entity

type FundValue struct {
	Id                  int    `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Code                string `gorm:"size:127;DEFAULT NULL" json:"fundcode"` //基金代码
	Name                string `gorm:"size:127;DEFAULT NULL" json:"name"`     //基金名称
	NetWorthDate        string `gorm:"size:127;DEFAULT NULL" json:"jzrq"`     //净值日期
	NetValueCurDay      string `gorm:"size:127;DEFAULT NULL" json:"dwjz"`     //当日净值
	EstimatedNetWorth   string `gorm:"size:127;DEFAULT NULL" json:"gsz"`      //估算净值
	EstimatedPercentage string `gorm:"size:127;DEFAULT NULL" json:"gszzl"`    //估算涨跌百分比
	EstimateTime        string `gorm:"size:127;DEFAULT NULL" json:"gztime"`   //估值时间
}
