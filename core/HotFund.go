package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
	"github.com/goinggo/mapstructure"
)

func (frc *FundRecCore) FetchHotFunds() {
	frc.mysqlDB.AutoMigrate(&entity.Fund{})

	rawRes := common.HttpGet(common.DXFundHotUrl)
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
