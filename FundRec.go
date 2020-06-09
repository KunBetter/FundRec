package main

import (
	"github.com/KunBetter/FundRec/doctorxiong"
	"github.com/KunBetter/FundRec/store"
	"github.com/KunBetter/FundRec/ttfund"
)

func main() {
	fs := &store.FundStore{}
	fs.Store()

	dxFund := &doctorxiong.DXFund{}
	dxFund.Test()

	ttFund := &ttfund.TTFund{}
	ttFund.Test()
}
