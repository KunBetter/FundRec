package core

import (
	"encoding/json"
	"fmt"
	"github.com/KunBetter/FundRec/common"
	"github.com/KunBetter/FundRec/entity"
	"github.com/KunBetter/FundRec/schedule"
)

type FundValueFetch struct {
	RecCore   *FundRecCore
	Scheduler *schedule.Scheduler
}

func (fv *FundValueFetch) Init() {
	fv.RecCore.DBRef().AutoMigrate(&entity.FundValue{})
}

func (fv *FundValueFetch) Start() {
	//engine start
	fv.Init()
	fv.Process()

	fv.Scheduler = schedule.NewScheduler(DetailFetchSpecTime, fv)
	fv.Scheduler.Start()
}

func (fv *FundValueFetch) Run() {
	fv.Process()
}

func (fv *FundValueFetch) Process() {
	var funds []entity.FundBaseInfo
	fv.RecCore.DBRef().Find(&funds)

	// TODO 并行
	var noValueFunds []string
	for i := 0; i < len(funds); i++ {
		fund := funds[i]
		code := fund.Code

		rawRes := common.HttpGet(fmt.Sprintf(common.TTFundValueUrl, code))
		if rawRes != "" {
			fundValue := &entity.FundValue{}
			body := rawRes[8 : len(rawRes)-2]
			if len(body) == 0 {
				noValueFunds = append(noValueFunds, code)
			} else {
				err := json.Unmarshal([]byte(body), &fundValue)
				if err != nil {
					fmt.Println("parse error: " + body + "," + code)
					noValueFunds = append(noValueFunds, code)
				} else {
					fv.RecCore.DBRef().Create(&fundValue)
				}
			}
		} else {
			noValueFunds = append(noValueFunds, code)
		}
	}
}
