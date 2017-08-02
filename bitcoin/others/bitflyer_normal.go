package others

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"testing"
	"strings"
	"fmt"
	"time"
)

type ExchangeItem struct {
	ID int `json:"id"`
	Side string `json:"side"`
	Price float64 `json:"price"`
	Size float64 `json:"size"`
	/*Exec_date string `json:"exec_date"`
	Buy_child_order_acceptance_id string`json:"buy_child_order_acceptance_id"`
	Sell_child_order_acceptance_id string `json:"sell_child_order_acceptance_id"`*/
}
type ExchangeItems struct {
	ExchangeItems []ExchangeItem
}

const (

 	URL  = "https://api.bitflyer.jp/v1/"
 	GET_BOARD  = "getboard"
	GET_EXECUTIONS = "getexecutions"
 	ETC_BTC  = "?product_code=ETH_BTC"
 	BTC_JPY  = "?product_code=BTC_JPY"
	COUNT_500 = "&count=500"
	COUNT_50 = "&count=50"
	COUNT_1000 = "&count=1000"

	BUY = "BUY"
	SELL = "SELL"
)

func TestName(t *testing.T) {

}
func getExecutions()  {
	for true  {
		time.Sleep(1 * time.Second)
		fmt.Println()
		resp ,err := http.Get(URL+GET_EXECUTIONS+BTC_JPY + COUNT_1000)

		if err != nil{
			panic(err)
		}
		exchangeItems := make([]ExchangeItem,0)
		jsonStr ,err := ioutil.ReadAll(resp.Body)
		if err != nil{
			panic(err)
		}
		json.Unmarshal(jsonStr,&exchangeItems);
		sumSell := float64(0)
		sumBuy := float64(0)
		sellCount := 0
		buyCount := 0
		for _,item := range exchangeItems{
			if (strings.Compare(item.Side,BUY) == 0){
				sumBuy += item.Price * item.Size
				buyCount++
			}
			if (strings.Compare(item.Side,SELL) == 0){
				sumSell += item.Price * item.Size
				sellCount++
			}
		}
		fmt.Printf("%s %f %s %d %s","卖出：" ,sumSell,"卖出人：",sellCount,"\n")

		fmt.Printf("%s %f %s %d ","买入：" ,sumBuy,"买入人：",buyCount)
	}

}