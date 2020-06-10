package doctorxiong

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
)

type DXFundHotResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    []entity.Fund  `json:"data"`
	Meta    map[string]int `json:"meta"`
}

func fetchDXFundHot() {
	rawRes := common.HttpGet(dxFundHotUrl)
	if rawRes == "" {
		return
	}

	fv := &DXFundHotResponse{}
	err := json.Unmarshal([]byte(rawRes), &fv)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(fv)
}
