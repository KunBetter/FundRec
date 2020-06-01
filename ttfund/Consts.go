package ttfund

const (
	/**
	 * 获取基金历史净值地址
	 * Historical Net Worth
	 * 参数：&code=150270&page=1&per=20&sdate=20200526&edate=20200526"
	 */
	hnwPrefix = "http://fund.eastmoney.com/f10/F10DataApi.aspx?type=lsjz"

	/**
	 * 获取基金列表
	 */
	fundsUrl = "http://fund.eastmoney.com/js/fundcode_search.js"

	/**
	 * 基金公司列表
	 */
	fundCompanyUrl = "http://fund.eastmoney.com/js/jjjz_gs.js"

	/**
	 * 基金估值
	 */
	fundValueUrl = "http://fundgz.1234567.com.cn/js/%s.js"
)
