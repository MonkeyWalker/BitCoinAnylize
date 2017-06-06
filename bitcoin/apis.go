package bitcoin

import (
	"github.com/gorilla/websocket"
	"log"
	"encoding/json"
	"time"
	"../globeconstant"
	"./myconstants"
//	"strings"
	"strings"
	"io"
	"bytes"
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
type SubSucess struct {
	ID string `json:"ID"`
	Status string `json:"status"`
	Subbed string `json:"subbed"`
	Ts int64 `json:"ts"`
}
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
func KeepPingPong(c *websocket.Conn)  {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	go func() {
		for true {
			var dat map[string]int64
			ReadJSON(c,&dat)

			if (dat["ping"] == 0){
				continue
			}
			dat["pong"] = dat["ping"]
			delete(dat,"ping")
			response,_  := json.Marshal(&dat)

			log.Println(string(response))
			err2 := c.WriteMessage(websocket.TextMessage, response)
			if err2 != nil {
				log.Println("write:", err2)
				return
			}

		}
	}()
}
func GetKlineInfo(c *websocket.Conn ,symbolID int, periodID int)(*SubResponse,error) {
	symbol := myconstants.GetSymbol(symbolID)
	period := myconstants.GetKlinePeriod(periodID)
	sub := MARKET + "." + symbol+ "." + KLINE+ "." + period
	subMsg := &SubMessage{Sub: sub, ID: ID}
	println(subMsg.Sub)
	c.WriteJSON(subMsg)
	subSuccess := SubSucess{}
	err := ReadJSON(c,&subSuccess)
	if err != nil{
		return nil,err
	}


	if strings.Compare(subSuccess.Status,globeconstant.OK) == 0{
		response := &SubResponse{}
		err := ReadJSON(c,response)
		return  response,err
	}else {
		return nil,nil
	}
}

func ReadJSON(c *websocket.Conn,v interface{}) error {
	//_, r, err := c.NextReader()
	_,m,err := c.ReadMessage()
	if err != nil {
		return err
	}
	depressData := depressGZIPStream(bytes.NewReader(m))
	err = json.Unmarshal(depressData,v)
	if err == io.EOF {
		// One value is expected in the message.
		err = io.ErrUnexpectedEOF
	}
	return err
}
