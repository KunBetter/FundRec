package core

type Scheduler struct {
}

/*
 * 调度策略
 * 1、基金公司、基金列表【天级】更新
 * 2、基金持仓情况，【天级】更新
 * 3、基金详情、基金净值，【增量】更新
 * 4、热门基金、排序基金列表，【天级】更新
 * 5、
 */

/*
	frc.FetchFundCompany()
	frc.FetchFundList()
	frc.FetchFundNetWorth("150270")
	frc.FetchFundValue("001186")
	frc.FetchFundPosition("001186")
	frc.FetchFund("202015")
	frc.FetchHotFunds()
	frc.FetchRankedFunds()
	frc.FetchDXFundDetail("003171")
*/
