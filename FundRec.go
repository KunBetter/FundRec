package main

import (
	"github.com/KunBetter/FundRec/cache"
	"github.com/KunBetter/FundRec/doctorxiong"
	"github.com/KunBetter/FundRec/store"
	"github.com/KunBetter/FundRec/ttfund"
)

func main() {
	bc := cache.BigCache{}
	bc.Test()

	fs := &store.FundStore{}
	fs.Store()

	dxFund := &doctorxiong.DXFund{}
	dxFund.Test()

	ttFund := &ttfund.TTFund{}
	ttFund.Test()
}
