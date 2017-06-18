package huobi

import (
	"github.com/gorilla/websocket"
	"../myconstants"
)

const MARKET  = "market"
const KLINE  = "kline"
const DEPTH = "depth"
const TRADE  = "trade"
const DETAIL  = "detail"
const ID   = "walker"

type SubMessage struct {
	Sub string  `json:"sub"`
	ID  string `json:"id"`
}
type SubSuccess struct {
	ID string `json:"ID"`
	Status string `json:"status"`
	Subbed string `json:"subbed"`
	Ts int64 `json:"ts"`
}
//订阅成功后的返回
type SubResponse struct {
	Ch string `json:"ch"`
	Ts int64 `json:"ts"`
	Tick struct {
		ID int `json:"ID"`
		Amount float64 `json:"amount"`
		Count int `json:"count"`
		Open float64 `json:"open"`
		Close float64 `json:"close"`
		Low float64 `json:"low"`
		High float64 `json:"high"`
		Vol float64 `json:"vol"`
	} `json:"tick"`
}

//订阅失败
type SubFail struct {
	ID string `json:"id"`
	Status string `json:"status"`
	ErrCode string `json:"err-code"`
	ErrMsg string `json:"err-jsonData"`
	Ts int64 `json:"ts"`
}

//取消订阅成功
type UnSubSuccess struct {
	ID string `json:"id"`
	Status string `json:"status"`
	Unsubbed string `json:"unsubbed"`
	Ts int64 `json:"ts"`
}


//订阅Line的消息
func SubKlineInfo(c *websocket.Conn ,symbolID int, periodID int) {
	symbol := myconstants.GetSymbol(symbolID)
	period := myconstants.GetKlinePeriod(periodID)
	sub := MARKET + "." + symbol+ "." + KLINE+ "." + period
	subMsg := &SubMessage{Sub: sub, ID: ID}
	println(subMsg.Sub)
	c.WriteJSON(subMsg)
}

