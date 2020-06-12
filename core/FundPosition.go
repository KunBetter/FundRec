package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
	"github.com/goinggo/mapstructure"
	"strconv"
	"strings"
)

func (frc *FundRecCore) FetchFundPosition(code string) {
	frc.mysqlDB.AutoMigrate(&entity.FundPosition{})
	frc.mysqlDB.AutoMigrate(&entity.FundStock{})

	rawRes := common.HttpGet(fmt.Sprintf(common.DXFundPositionUrl, code))
	if rawRes == "" {
		return
	}

	var buf map[string]interface{}
	err := json.Unmarshal([]byte(rawRes), &buf)
	if err != nil {
		fmt.Println("some error")
	}

	data := buf["data"].(map[string]interface{})
	stockList := data["stockList"].([]interface{})
	delete(data, "stockList")

	var fundStockIds []string
	//stock store
	for i := 0; i < len(stockList); i++ {
		sbuf := stockList[i].([]interface{})
		fundStock := entity.FundStock{
			Code:  sbuf[0].(string),
			Name:  sbuf[1].(string),
			Ratio: sbuf[2].(string),
			Num:   sbuf[3].(string),
			Cash:  sbuf[4].(string),
		}
		fundStock.FundCode = code
		frc.DBRef().Create(&fundStock)
		fundStockIds = append(fundStockIds, strconv.Itoa(fundStock.Id))
	}

	var fundPosition entity.FundPosition
	err = mapstructure.Decode(data, &fundPosition)
	if err != nil {
		fmt.Println("some error")
	}
	fundPosition.Code = code
	fundPosition.FundStockIDs = strings.Join(fundStockIds, ",")
	frc.DBRef().Create(&fundPosition)
}
