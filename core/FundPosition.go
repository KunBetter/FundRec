package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
)

type DXFPosResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    entity.FundPosition `json:"data"`
	Meta    string              `json:"meta"`
}

func (frc *FundRecCore) FetchDXFundPosition(code string) {
	rawRes := common.HttpGet(fmt.Sprintf(common.DXFundPositionUrl, code))
	if rawRes == "" {
		return
	}

	fv := &DXFPosResponse{}
	err := json.Unmarshal([]byte(rawRes), &fv)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(fv)
}
