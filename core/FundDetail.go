package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
	"github.com/goinggo/mapstructure"
	"strconv"
)

func (frc *FundRecCore) FetchDXFundDetail(code string) {
	frc.mysqlDB.AutoMigrate(&entity.FundIncome{})
	frc.mysqlDB.AutoMigrate(&entity.FundDetail{})
	frc.mysqlDB.AutoMigrate(&entity.FundNetWorth{})

	rawRes := common.HttpGet(fmt.Sprintf(common.DXFundDetailUrl, code))
	if rawRes == "" {
		return
	}

	var buf map[string]interface{}
	err := json.Unmarshal([]byte(rawRes), &buf)
	if err != nil {
		fmt.Println("some error")
	}

	data := buf["data"].(map[string]interface{})

	//netWorthData
	if value, ok := data["netWorthData"]; ok {
		netWorthData := value.([]interface{})
		delete(data, "netWorthData")
		for i := 0; i < len(netWorthData); i++ {
			fundNetWorthData := netWorthData[i].([]interface{})
			fundNetWorth := entity.FundNetWorth{
				Date:            fundNetWorthData[0].(string),
				Unit:            float32(fundNetWorthData[1].(float64)),
				DailyGrowthRate: strconv.FormatFloat(fundNetWorthData[2].(float64), 'f', 4, 32),
				Dividend:        fundNetWorthData[3].(string),
				Code:            code,
			}
			frc.DBRef().Create(&fundNetWorth)
		}
	}

	//millionCopiesIncomeData
	if value, ok := data["millionCopiesIncomeData"]; ok {
		millionCopiesIncomeData := value.([]interface{})
		delete(data, "millionCopiesIncomeData")
		for i := 0; i < len(millionCopiesIncomeData); i++ {
			fundMillionCopiesIncomeData := millionCopiesIncomeData[i].([]interface{})
			fundIncome := entity.FundIncome{
				Date:   fundMillionCopiesIncomeData[0].(string),
				Income: float32(fundMillionCopiesIncomeData[1].(float64)),
				Type:   1,
				Code:   code,
			}
			frc.DBRef().Create(&fundIncome)
		}
	}

	//sevenDaysYearIncomeData
	if value, ok := data["sevenDaysYearIncomeData"]; ok {
		sevenDaysYearIncomeData := value.([]interface{})
		delete(data, "sevenDaysYearIncomeData")
		for i := 0; i < len(sevenDaysYearIncomeData); i++ {
			fundSevenDaysYearIncomeData := sevenDaysYearIncomeData[i].([]interface{})
			fundIncome := entity.FundIncome{
				Date:   fundSevenDaysYearIncomeData[0].(string),
				Income: float32(fundSevenDaysYearIncomeData[1].(float64)),
				Type:   2,
				Code:   code,
			}
			frc.DBRef().Create(&fundIncome)
		}
	}

	var fundDetail entity.FundDetail
	err = mapstructure.Decode(data, &fundDetail)
	if err != nil {
		fmt.Println("some error")
	}
	frc.DBRef().Create(&fundDetail)
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
