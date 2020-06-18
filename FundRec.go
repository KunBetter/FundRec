package main

import (
	"github.com/KunBetter/FundRec/core"
)

func main() {
	frc := &core.FundRecCore{}
	frc.Init()
	frc.FundDataFetch()
	frc.Router()
}
