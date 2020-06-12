package entity

//基金持仓
type FundPosition struct {
	Id           int    `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Code         string `gorm:"size:127;DEFAULT NULL" json:"code"`
	Title        string `gorm:"size:127;DEFAULT NULL" json:"title"`        //标题信息
	Date         string `gorm:"size:127;DEFAULT NULL" json:"date"`         //截至时间
	Stock        string `gorm:"size:127;DEFAULT NULL" json:"stock"`        //股票占比
	Bond         string `gorm:"size:127;DEFAULT NULL" json:"bond"`         //债券占比
	Cash         string `gorm:"size:127;DEFAULT NULL" json:"cash"`         //现金占比
	Total        string `gorm:"size:127;DEFAULT NULL" json:"total"`        //总净资产(亿元)
	FundStockIDs string `gorm:"size:127;DEFAULT NULL" json:"fundStockIDs"` //基金股票详情列表IDs
}

type FundStock struct {
	Id       int    `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	FundCode string `gorm:"size:127;DEFAULT NULL" json:"fundCode"` //基金代码
	Code     string `gorm:"size:127;DEFAULT NULL" json:"code"`     //股票代码
	Name     string `gorm:"size:127;DEFAULT NULL" json:"name"`     //名称
	Ratio    string `gorm:"size:127;DEFAULT NULL" json:"ratio"`    //占比
	Num      string `gorm:"size:127;DEFAULT NULL" json:"num"`      //持有股数(万股)
	Cash     string `gorm:"size:127;DEFAULT NULL" json:"cash"`     //持有金额(万元)
}
