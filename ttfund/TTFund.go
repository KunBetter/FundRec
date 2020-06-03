package ttfund

import (
	"fmt"
)

type TTFund struct {
}

func genHistoricalNetWorthUrl(fundCode string, pageIdx int, pageSize int, sDate string, eDate string) string {
	return fmt.Sprintf(hnwPrefix+"&code=%s&page=%d&per=%d&sdate=%s&edate=%s", fundCode, pageIdx, pageSize, sDate, eDate)
}

func (ttFund *TTFund) Test() {
	fetchFundValue("001186")
	fetchFundCompany()
	fetchFundList()
	fecthFundNetWorth("150270")
}
