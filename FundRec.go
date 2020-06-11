package main

import (
	"github.com/KunBetter/FundRec/core"
	"github.com/KunBetter/FundRec/store"
)

func main() {
	frc := &core.FundRecCore{}
	frc.Init()
	frc.Run()

	fs := &store.FundStore{}
	fs.Store()
}
