package core

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

type DXFResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []entity.Fund `json:"data"`
	Meta    string        `json:"meta"`
}

func (frc *FundRecCore) FetchDXFund(code string) {
	rawRes := common.HttpGet(fmt.Sprintf(common.DXFundUrl, code))
	if rawRes == "" {
		return
	}

	fv := &DXFResponse{}
	err := json.Unmarshal([]byte(rawRes), &fv)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(fv)
}
