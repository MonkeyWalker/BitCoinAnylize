package bitcoin

import (
	"log"
	"./myconstants"
	"github.com/gorilla/websocket"
	"io"
	"compress/gzip"
	"io/ioutil"
	"os"
	"os/signal"
	"time"

)


//将接收到的GZIP压缩数据解压
func depressGZIPStream(datasSlice io.Reader)  []byte{
	r,err := gzip.NewReader(datasSlice)
	defer r.Close()
	recover()
	if err != nil{
		panic(err)
	}
	undatas, _ := ioutil.ReadAll(r)
	return undatas
}

func HuoBiMarket() {
	recover()
	u := myconstants.WS_BTC
	log.Printf("connecting to %s", u)
	//建立链接
	c, _, err := websocket.DefaultDialer.Dial(u, nil)
	defer c.Close()
	go ReadJSON(c)
	go DealWithResponse()
	if err != nil {
		log.Fatal("dial:", err)
	}
	time.After(2 * time.Second)
	SubKlineInfo(c,myconstants.BTC_SYMBOL,myconstants.MIN_5)
	if err != nil{
		panic(err)
	}
	waitInterrupt()
}

//等待外部传入的中断信号
func waitInterrupt()  {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for true {
		c := <-interrupt
		log.Println(c)
	}
	// To cleanly close a connection, a client should send a close
	// frame and wait for the server to close the connection.
	return
}