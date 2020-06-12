package common

const (
	FundCompanyCache = "fundCompany"
)

const (
	//******************************熊博士数据源****************************************
	// https://www.doctorxiong.club
	/**
	 * GetV1Fund - 获取基金基础信息
	 */
	DXFundUrl = "https://api.doctorxiong.club/v1/fund?code=%s"

	/**
	 * GetV1FundPosition - 获取基金持仓详情
	 */
	DXFundPositionUrl = "https://api.doctorxiong.club/v1/fund/position?code=%s"

	/**
	 * PostV1FundRank - 获取基金排行
	 */
	DXFundRankUrl = "https://api.doctorxiong.club/v1/fund/rank"

	/**
	 * GetV1FundDetail - 获取基金详情
	 */
	DXFundDetailUrl = "https://api.doctorxiong.club/v1/fund/detail?code=%s"

	/**
	 * GetV1FundHot - 获取热门基金
	 */
	DXFundHotUrl = "https://api.doctorxiong.club/v1/fund/hot"

	//***************************天天基金数据源*****************************************
	/**
	 * 获取基金历史净值地址
	 * Historical Net Worth
	 * 参数：&code=150270&page=1&per=20&sdate=20200526&edate=20200526"
	 */
	HNWPrefix = "http://fund.eastmoney.com/f10/F10DataApi.aspx?type=lsjz"

	/**
	 * 获取基金列表
	 */
	TTFundsListUrl = "http://fund.eastmoney.com/js/fundcode_search.js"

	/**
	 * 基金公司列表
	 */
	TTFundCompanyUrl = "http://fund.eastmoney.com/js/jjjz_gs.js"

	/**
	 * 基金估值
	 */
	TTFundValueUrl = "http://fundgz.1234567.com.cn/js/%s.js"
)
