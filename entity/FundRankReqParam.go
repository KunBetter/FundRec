package entity

type FundRankReqParam struct {
	/*
		基金种类(可以接受多个参数)
		預設值: 所有类型
		Allowed values: {"股票型", "混合型", "债券型", "指数型", "QDII", "FOF"}
	*/
	FundType []string `json:"fundType"`
	/* 	dayGrowth日涨幅
	  	lastWeek最近一周,
		lastMonth最近一个月涨幅排序；
		預設值: dayGrowth ；
		Allowed values: {"dayGrowth", "lastWeekGrowth", "lastMonthGrowth", "lastThreeMonthsGrowth",
			"lastSixMonthsGrowth", "thisYearGrowth", "lastYearGrowth", "lastTwoYearsGrowth", "lastThreeYearsGrowth"}
	*/
	Sort string `json:"sort"`
	/*	基金公司,选择"易方达"就只返回易方达的基金(接受多个参数) ；
		預設值: 所有基金公司 ；
		Allowed values: {"华夏", "嘉实", "易方达", "南方", "中银", "广发", "工银", "博时", "华安", "汇添富"}
	*/
	FundCompany    []string `json:"fundCompany"`
	FundGrade      []string `json:"fundGrade"`      //基金认证等级,也支持一,二,三,四星(非必需) ；預設值: 无认证等级要求；Allowed values: {"上证五星", "招商五星", "济安五星"}
	CreatTimeLimit int      `json:"creatTimeLimit"` //基金成立时间限制1:小于一年》2:小于两年.依次类推(非必需) ；預設值: 1
	FundScale      int      `json:"fundScale"`      //基金规模单位亿,10表示10亿以下,100表示100亿以下,1001表示100亿以上(非必需) ；Allowed values: [10, 100, 1001]
	Asc            int      `json:"asc"`            //排序方式0:降序:1升序；預設值: 0 ；Allowed values: {0，1}
	PageIndex      int      `json:"pageIndex"`      //页码；預設值: 1
	PageSize       int      `json:"pageSize"`       //每页显示的数量；預設值: 10
}
