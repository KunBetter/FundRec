package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
	"github.com/goinggo/mapstructure"
	"strconv"
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

		netWorthStr := fBuf["netWorth"].(string)
		delete(fBuf, "netWorth")
		value, err := strconv.ParseFloat(netWorthStr, 32)
		if err != nil {
			fmt.Println("some error")
		}
		netWorth := float32(value)

		var fund entity.Fund
		err = mapstructure.Decode(fBuf, &fund)
		if err != nil {
			fmt.Println("some error")
		}

		fund.NetWorth = netWorth
		frc.DBRef().Create(&fund)
	}
}
