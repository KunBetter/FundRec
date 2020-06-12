package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
)

func (frc *FundRecCore) FetchFundValue(code string) {
	frc.mysqlDB.AutoMigrate(&entity.FundValue{})

	rawRes := common.HttpGet(fmt.Sprintf(common.TTFundValueUrl, code))
	if rawRes == "" {
		return
	}

	fundValue := &entity.FundValue{}
	err := json.Unmarshal([]byte(rawRes[8:len(rawRes)-2]), &fundValue)
	if err != nil {
		fmt.Println("some error")
	}
	frc.DBRef().Create(&fundValue)
}
