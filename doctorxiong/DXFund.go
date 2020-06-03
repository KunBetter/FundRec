package doctorxiong

/**
 * https://www.doctorxiong.club
 */

const (
	/**
	 * GetV1Fund - 获取基金基础信息
	 */
	dxFundUrl = "https://api.doctorxiong.club/v1/fund?code=%s"

	/**
	 * GetV1FundPosition - 获取基金持仓详情
	 */
	dxFundPosUrl = "https://api.doctorxiong.club/v1/fund/position?code=%s"

	/**
	 * PostV1FundRank - 获取基金排行
	 */
	dxFundRankUrl = "https://api.doctorxiong.club/v1/fund/rank"

	/**
	 * GetV1FundDetail - 获取基金详情
	 */
	dxFundDetailUrl = "https://api.doctorxiong.club/v1/fund/detail?code=%s"

	/**
	 * GetV1FundHot - 获取热门基金
	 */
	dxFundHotUrl = "https://api.doctorxiong.club/v1/fund/hot"
)

type DXFund struct {
}

func (dxf *DXFund) Test() {
	fetchDXFundRank()
	fetchDXFundHot()
	fetchDXFundDetail("000001")
	fetchDXFund("202015")
	fetchDXFundPosition("202015")
}
