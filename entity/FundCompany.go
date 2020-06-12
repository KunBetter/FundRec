package entity

type FundCompany struct {
	Id   int    `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	Code string `gorm:"size:127;DEFAULT NULL" json:"code"` //基金公司代码
	Name string `gorm:"size:127;DEFAULT NULL" json:"name"` //基金公司名称中文
}
