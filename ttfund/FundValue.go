package ttfund

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
)

func fetchFundValue(code string) {
	rawRes := common.HttpGet(fmt.Sprintf(fundValueUrl, code))
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
