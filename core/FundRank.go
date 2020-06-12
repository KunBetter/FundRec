package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
	"github.com/goinggo/mapstructure"
)

func (frc *FundRecCore) FetchRankedFunds() {
	frp := entity.FundRankReqParam{
		FundType:    []string{"指数型", "股票型", "债券型", "QDII", "FOF", "混合型"},
		Sort:        "lastWeekGrowth",
		FundCompany: []string{"嘉实", "易方达", "博时", "华夏", "中银", "工银", "广发", "南方", "华安", "汇添富"},
		FundGrade:   []string{"上证五星"},
	}

	b, _ := json.Marshal(frp)
	rawRes := common.HttpPost(common.DXFundRankUrl, "application/json", string(b))
	if rawRes == "" {
		return
	}

	var buf map[string]interface{}
	err := json.Unmarshal([]byte(rawRes), &buf)
	if err != nil {
		fmt.Println("some error")
	}

	data := buf["data"].([]interface{})

	for i := 0; i < len(data); i++ {
		fBuf := data[i].(map[string]interface{})

		netWorth := common.Str2Float32(fBuf["netWorth"].(string))
		delete(fBuf, "netWorth")

		var fund entity.Fund
		err = mapstructure.Decode(fBuf, &fund)
		if err != nil {
			fmt.Println("some error")
		}

		fund.NetWorth = netWorth
		frc.DBRef().Create(&fund)
	}
}
