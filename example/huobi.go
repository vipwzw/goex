package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/huobi"
)

var dsn = flag.String("dsn", "root:wzw123456@/trade", "mysql data source name")

func main() {
	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		panic(err)
	}
	ws := huobi.NewHbdmWs() //huobi期货
	//设置回调函数
	ws.SetCallbacks(func(ticker *goex.FutureTicker) {
		//printJSON(ticker)
	}, func(depth *goex.Depth) {
		log.Println(depth)
	}, func(trade *goex.Trade, contract string) {
		saveToDB(db, trade)
		if trade.Amount > 200 {
			log.Println(trade)
		}
	})
	//订阅行情
	ws.SubscribeTrade(goex.BTC_USDT, goex.QUARTER_CONTRACT)
	//ws.SubscribeDepth(goex.BTC_USDT, goex.QUARTER_CONTRACT, 5)
	//ws.SubscribeTicker(goex.BTC_USDT, goex.QUARTER_CONTRACT)
	select {}
}

func printJSON(v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	log.Println(string(data))
}

func saveToDB(db *sql.DB, trade *goex.Trade) {
	q := fmt.Sprintf(
		"insert into trade.huobi_quarter_btcusdt(tid, type, amount, price, time_ms) values ('%d', '%d', '%f', '%f', '%d')", trade.Tid, trade.Type, trade.Amount, trade.Price, trade.Date)
	tryInsert(db, q)
}

func tryInsert(db *sql.DB, q string) {
	_, err := db.Exec(q)
	if err != nil {
		log.Println("exec failed:", err, ", sql:", q)
	}
}
