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
	trends []float32
};

const URL  = "https://api.bitflyer.jp/v1/getboard"
const ETC_BTC  = "?product_code=ETH_BTC"
const BTC_JPY  = "?product_code=BTC_JPY"



func (btc *btc) Init()  {
	btc.values = make([]float32,1)
	btc.trends = make([]float32,1)
}

func (btc *btc) PushTrend(trend float32) {
	btc.trends = append(btc.trends, trend)
}

func (btc *btc) PushValue(value float32)  {
	btc.values = append(btc.values, value)
}

func (btc *btc) PopTrend()   {
	if (len(btc.trends) < 1){
		return
	}
	btc.trends = btc.trends[1:]
}


func (btc *btc) ClearTrend()   {
	btc.trends = nil;
}

func (btc *btc) PopValue()  {
	if (len(btc.values) < 1){
		return
	}
	btc.values = btc.values[1:]
}



func Bitflyer() {
	btc := &btc{}
	btc.Init();
	cacluteBTCTrend(btc)
}

//TODO:；利用实时汇率对日元进行计算
func getExchangeRate()  {}

func cacluteBTCTrend(btc *btc)  {
	for true{
		//每10秒拉下来 一次数据
		time.Sleep(10 * time.Second)
		newest := getBoard(URL + BTC_JPY)
		fmt.Printf("%s：%f","get the value : " , newest)
		if (len(btc.values) > 100){
			btc.PopValue()
		}
		btc.PushValue(newest)
		//先清空 趋势
		btc.ClearTrend();
		for  index := 0; index < len(btc.values) - 1;index++ {
			btc.PushTrend(( newest - btc.values[index]) / newest)
		}
		fmt.Println(btc.trends)
		fmt.Println(btc.values)
	}

}
func getBoard(url string) float32{

	re,error := http.Get(url);
	if (error != nil){
		panic(error)
	}
	defer re.Body.Close();
	body, _ := ioutil.ReadAll(re.Body)
	marketBoard := 	&MarketBoard{}
	json.Unmarshal(body,marketBoard)
	println(marketBoard.Mid_Price)
	return marketBoard.Mid_Price * jpyExchange
}