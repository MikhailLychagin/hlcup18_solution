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
		Accounts: make(map[int16]*models.AccountEntry),
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
	// var acc models.AccountEntry
	acc := &models.AccountEntry{}
	err := acc.UnmarshalJSON(ctx.PostBody())
	if err != nil {
		panic(err)
	}
	dm.AccountAdd(acc)
	fmt.Print(acc)
	accStr, _ := json.Marshal(acc)
	ctx.Write(accStr)
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}
