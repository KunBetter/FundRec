package main

import (
	"github.com/KunBetter/FundRec/doctorxiong"
	"github.com/KunBetter/FundRec/ttfund"
)

func main() {
	dxFund := &doctorxiong.DXFund{}
	dxFund.Test()

	ttFund := &ttfund.TTFund{}
	ttFund.Test()
}
