package entity

type FundBaseInfo struct {
	Code        string `json:"code"`        //基金代码
	Name        string `json:"name"`        //基金名称中文
	Type        string `json:"type"`        //基金类型
	SingleSpell string `json:"singleSpell"` //基金名称单拼
	AllSpell    string `json:"allSpell"`    //基金名称全拼
}

type Fund struct {
	FundBaseInfo
	NetWorth                float32         `json:"netWorth"`                //当前基金单位净值
	ExpectWorth             float32         `json:"expectWorth"`             //当前基金单位净值估算
	TotalWorth              float32         `json:"totalWorth"`              //当前基金累计净值
	ExpectGrowth            string          `json:"expectGrowth"`            //当前基金单位净值估算日涨幅,单位为百分比
	DayGrowth               string          `json:"dayGrowth"`               //单位净值日涨幅,单位为百分比
	LastWeekGrowth          string          `json:"lastWeekGrowth"`          //单位净值周涨幅,单位为百分比
	LastMonthGrowth         string          `json:"lastMonthGrowth"`         //单位净值月涨幅,单位为百分比
	LastThreeMonthsGrowth   string          `json:"lastThreeMonthsGrowth"`   //单位净值三月涨幅,单位为百分比
	LastSixMonthsGrowth     string          `json:"lastSixMonthsGrowth"`     //单位净值六月涨幅,单位为百分比
	LastYearGrowth          string          `json:"lastYearGrowth"`          //单位净值年涨幅,单位为百分比
	ThisYearGrowth          string          `json:"thisYearGrowth"`          //今年的涨幅,单位为百分比
	BuyMin                  float32         `json:"buyMin"`                  //起购额度
	BuySourceRate           float32         `json:"buySourceRate"`           //原始买入费率,单位为百分比
	BuyRate                 float32         `json:"buyRate"`                 //当前买入费率,单位为百分比
	Manager                 string          `json:"manager"`                 //基金经理
	FundScale               string          `json:"fundScale"`               //基金规模及日期,日期为最后一次规模变动的日期
	WorthDate               string          `json:"worthDate"`               //净值更新日期,日期格式为yy-MM-dd HH:mm.2019-06-27 15:00代表当天下午3点
	ExpectWorthDate         string          `json:"expectWorthDate"`         //净值估算更新日期,,日期格式为yy-MM-dd HH:mm.2019-06-27 15:00代表当天下午3点
	NetWorthData            [][]interface{} `json:"netWorthData"`            //历史净值信息["2001-12-18" , 1 , 0 , ""]依次表示:日期; 单位净值; 净值涨幅; 每份分红.
	MillionCopiesIncome     float32         `json:"millionCopiesIncome"`     //每万分收益(货币基金)
	MillionCopiesIncomeData [][]interface{} `json:"millionCopiesIncomeData"` //历史万每分收益信息(货币基金)["2016-09-23",0.4773]依次表示:日期; 每万分收益.
	MillionCopiesIncomeDate string          `json:"millionCopiesIncomeDate"` //七日年化收益更新日期(货币基金)
	SevenDaysYearIncome     float32         `json:"sevenDaysYearIncome"`     //七日年化收益(货币基金)
	SevenDaysYearIncomeData [][]interface{} `json:"sevenDaysYearIncomeData"` //历史七日年华收益信息(货币基金)["2016-09-23",2.131]依次表示:日期; 七日年化收益.
}

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

type FundValue struct {
	Code                string `json:"fundcode"` //基金代码
	Name                string `json:"name"`     //基金名称
	NetWorthDate        string `json:"jzrq"`     //净值日期
	NetValueCurDay      string `json:"dwjz"`     //当日净值
	EstimatedNetWorth   string `json:"gsz"`      //估算净值
	EstimatedPercentage string `json:"gszzl"`    //估算涨跌百分比
	EstimateTime        string `json:"gztime"`   //估值时间
}
