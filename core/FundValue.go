package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
)

func (frc *FundRecCore) FetchFundValue(code string) { //001186
	rawRes := common.HttpGet(fmt.Sprintf(common.TTFundValueUrl, code))
	if rawRes == "" {
		return
	}

	fv := &entity.FundValue{}
	err := json.Unmarshal([]byte(rawRes[8:len(rawRes)-2]), &fv)
	if err != nil {
		fmt.Println("some error")
	}
	fmt.Println(fv)
}
