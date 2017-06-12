package main

import (
	//"./bitcoin"
	"./globeutils"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
	}()
	//bitcoin.HuoBiMarket()
	//testSoSoBTC()
	go connnectWS("wss://io.sosobtc.com/sosobtc.io/?chnl=okcoin&EIO=3&transport=websocket&sid=ETFdY92Cf1LA5uwiApH4")
	connnectWS("wss://io.sosobtc.com/sosobtc.io/?EIO=3&transport=websocket")

}
func connnectWS(u string)  {
	log.Printf("connecting to %s", u)
	headers := http.Header{}
	globeutils.ReadLine("./headers.txt", func(line string) {
		strs := strings.Split(line, ",")
		if len(strs) > 1 {
			headers.Set(strs[0], strs[1])
		}
	})

	c1, res, err2 := websocket.DefaultDialer.Dial(u, headers)
	if res != nil{
		println(res.StatusCode)
	}
	if err2 != nil {
		panic(err2)
	}
	for true {
		time.Sleep(1 * time.Second)
		_, m, _ := c1.ReadMessage()
		//if (len(m) > 0){
			println(string(m))
		//}
	}

}

