package others

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	"fmt"
)

/**
 * bitflyer 是日本最大的数字货币交易平台,日本又是当前的数字货币第三大市场
 * 实时监控 其异常波动情况,在
 */
var jpyExchange float32 = 0.06141
type Bids struct {
 	Price float32 `price`
	Size  float32 `size`
}

type MarketBoard struct {
	Mid_Price float32 `mid_price`
	Bids []Bids `bids`
}

type btc struct {
	// 比特币的价格,每5秒请求一次
	values []float32

	//过去的数据相对于最新数据的差值！后的比例
	valueTrends []float32

	// 交易的数量
	amounts []float32

	// 交易的数量变化
	amountTrends []float32

};


func (btc *btc) Init()  {
	btc.values = make([]float32,0)
	btc.valueTrends = make([]float32,0)
	btc.amounts = make([]float32,0)
	btc.valueTrends = make([]float32,0)
}

func (btc *btc) PushTrend(trend float32) {
	btc.valueTrends = append(btc.valueTrends, trend)
}

func (btc *btc) PushValue(value float32)  {
	btc.values = append(btc.values, value)
}

func (btc *btc) PopTrend()   {
	if (len(btc.valueTrends) < 1){
		return
	}
	btc.valueTrends = btc.valueTrends[1:]
}


func (btc *btc) ClearTrend()   {
	btc.valueTrends = nil;
}

func (btc *btc) PopValue()  {
	if (len(btc.values) < 1){
		return
	}
	btc.values = btc.values[1:]
}



func Bitflyer() {
	/**/
	//for true{

		fmt.Println()
		go getExecutions();
	//}
	btc := &btc{}
	btc.Init();
	cacluteBTCTrend(btc)

}

//TODO:；利用实时汇率对日元进行计算
func getExchangeRate()  {}

func cacluteBTCTrend(btc *btc)  {
	for true{
		//每5秒拉下来 一次数据
		time.Sleep(5 * time.Second)
		newest := getBoard(URL + GET_BOARD + BTC_JPY)

		if (len(btc.values) > 100){
			btc.PopValue()
		}
		btc.PushValue(newest.Mid_Price)
		//先清空 趋势
		btc.ClearTrend()
		for  index := 0; index < len(btc.values) - 1;index++ {
			trend := (newest.Mid_Price  - btc.values[index] )/ newest.Mid_Price
			btc.PushTrend(trend)
		}
		fmt.Println()
		fmt.Println(btc.values)

		fmt.Println()
		fmt.Println(btc.valueTrends)
	}

}
func getBoard(url string) *MarketBoard{

	re,error := http.Get(url);
	if (error != nil){
		panic(error)
	}
	defer re.Body.Close();
	body, _ := ioutil.ReadAll(re.Body)
	marketBoard := 	&MarketBoard{}
	json.Unmarshal(body,marketBoard)

	return marketBoard
}