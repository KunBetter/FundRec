package doctorxiong

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
)

type DXFundDetailResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    entity.Fund `json:"data"`
	Meta    string      `json:"meta"`
}

func fetchDXFundDetail(code string) {
	rawRes := common.HttpGet(fmt.Sprintf(dxFundDetailUrl, code))
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
