package core

const (
	BaseFetchSpecTime   = "0 0 8 * ?"
	DetailFetchSpecTime = "0 0 9 * ?"
)

type Fetch interface {
	Init()
	Process()
}

/*
	frc.FetchFundNetWorth("150270")
	frc.FetchFundPosition("001186")
	frc.FetchFund("202015")
	frc.FetchDXFundDetail("003171")

	frc.FetchHotFunds()
	frc.FetchRankedFunds()
*/
