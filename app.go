package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/mhc18_go/models"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

var (
	addr = flag.String("addr", ":28080", "TCP address to listen to")
	dm   = &models.DataManager{
		Accounts:     make(map[uint16]*models.AccountEntry),
		CountryToId:  make(models.ReprToIdMap, 1024),
		IdToCountry:  make([]*string, 1024),
		CityToId:     make(models.ReprToIdMap, 1024),
		IdToCity:     make([]*string, 1024),
		InterestToId: make(models.ReprToIdMap, 1024),
		IdToInterest: make([]*string, 1024),
		CurrentDate:  1547753597,
		IndexManager: *models.BuildDefaultIndexManager(),
	}
)

func main() {
	router := routing.New()

	router.Post("/accounts/new/", accountNew)

	if err := fasthttp.ListenAndServe(*addr, router.HandleRequest); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func accountNew(ctx *routing.Context) error {
	// var acc models.AccountFormEntry
	accForm := &models.AccountFormEntry{}
	err := accForm.UnmarshalJSON(ctx.PostBody())
	if err != nil {
		fmt.Println(err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return nil
	}
	fmt.Println(accForm)
	acc, err := dm.FormToAccount(accForm)
	if err != nil {
		fmt.Println(err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return nil
	}
	fmt.Println(acc)
	dm.AccountAdd(acc)
	accStr, _ := json.Marshal(acc)
	ctx.Write(accStr)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}
