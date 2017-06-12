package bitcoin

import (
	"strings"
	"github.com/gorilla/websocket"
	"bytes"
	"encoding/json"
	"log"
)

type Response struct {
	SubResponse *SubResponse
	SubFail     *SubFail
}
func handlePingPong(c *websocket.Conn,jsonData string){
	dat := make(map[string]int64)
	err  := json.Unmarshal([]byte(jsonData),&dat)
	if (err != nil){
		panic(err)
	}
	dat["pong"] = dat["ping"]
	delete(dat,"ping") //删除键 ping对应的值

	response,err := json.Marshal(&dat)
	if (err != nil){
		panic(err)
	}

	err2 := c.WriteMessage(websocket.TextMessage, response)
	if err2 != nil {
		log.Println("write:", err2)
		return
	}
}

//处理订阅返回的信息
func handleSubResponse(jsonData string) *SubResponse {
	subResponse := &SubResponse{}
	err := json.Unmarshal([]byte(jsonData),&subResponse)
	if err != nil{
		panic(err)
	}
	return subResponse
}
//

func handleSubFail(jsonData string) *SubFail {
	subFail := &SubFail{}
	unMarshalJson(jsonData,subFail)
	return subFail
}

func unMarshalJson(jsonData string,v interface{})  {
	err := json.Unmarshal([]byte(jsonData),v)
	if err != nil{
		panic(err)
	}
}

var blockList chan *Response
func ReadJSON(c *websocket.Conn) {
	blockList = make(chan *Response,5)
	for true {
		_,m,err := c.ReadMessage()
		if err != nil {
			panic(err)
		}
		depressData := depressGZIPStream(bytes.NewReader(m))
		jsonData := string(depressData)
		//println(jsonData)
		if strings.Contains(jsonData,"ping"){
			handlePingPong(c,jsonData)
		}else if strings.Contains(jsonData,"tick") {
			r := Response{SubResponse: handleSubResponse(jsonData)}
			blockList <- &r
		}else if strings.Contains(jsonData, "err-code"){
			r := Response{SubFail:handleSubFail(jsonData)}
			blockList <- &r
		}
	}
}

//处理返回的数据
func DealWithResponse(){
	for true {
		/*res := <- blockList
		println(res.SubResponse.Tick.Low)*/
	}

}
