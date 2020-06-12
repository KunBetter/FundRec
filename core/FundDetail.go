package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
	"github.com/goinggo/mapstructure"
)

type DXFundDetailResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    entity.FundDetail `json:"data"`
	Meta    string            `json:"meta"`
}

func (frc *FundRecCore) FetchDXFundDetail(code string) {
	rawRes := common.HttpGet(fmt.Sprintf(common.DXFundDetailUrl, code))
	if rawRes == "" {
		return
	}

	fv := &DXFundDetailResponse{}
	err := json.Unmarshal([]byte(rawRes), &fv)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(fv)
}

func (frc *FundRecCore) FetchFund(code string) {
	frc.mysqlDB.AutoMigrate(&entity.Fund{})

	rawRes := common.HttpGet(fmt.Sprintf(common.DXFundUrl, code))
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

		var fund entity.Fund
		err = mapstructure.Decode(fBuf, &fund)
		if err != nil {
			fmt.Println("some error")
		}

		frc.DBRef().Create(&fund)
	}
}
