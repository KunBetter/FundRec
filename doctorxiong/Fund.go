package doctorxiong

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
)

type DXFResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []entity.Fund `json:"data"`
	Meta    string        `json:"meta"`
}

func fetchDXFund(code string) {
	rawRes := common.HttpGet(fmt.Sprintf(dxFundUrl, code))
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
