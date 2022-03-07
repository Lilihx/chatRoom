/*
 * @Description: Add file description
 * @Author: lilihx@github.com
 * @Date: 2022-03-04 16:43:48
 * @LastEditTime: 2022-03-07 19:14:48
 * @LastEditors: lilihx@github.com
 */
package wss

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/lilihx/chatRoom/common/discover"
)

type WServer struct {
}

func NewServer() *WServer {
	return &WServer{}
}

func (ws *WServer) InitWebSocket() error {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.serve(w, r)
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write(nil)
	})

	consulClient, err := discover.NewKitDiscoverClient("127.0.0.1", 8500)
	if err != nil {
		fmt.Println("Init consul client err:" + err.Error())
		return err
	}
	consulClient.Register("wss", "123456", "/health", "127.0.0.1", 7000, nil)
	defer consulClient.DeRegister("123456")

	err = http.ListenAndServe("127.0.0.1:7000", nil)

	if err != nil {
		fmt.Println("Init websocket err:" + err.Error())
	}
	return err
}

func (ws *WServer) serve(w http.ResponseWriter, r *http.Request) {
	var upGrader = websocket.Upgrader{
		ReadBufferSize:  512,
		WriteBufferSize: 512,
	}
	upGrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("init websocket err")
		return
	}
	go ws.readMsg(conn)
}

func (ws *WServer) readMsg(conn *websocket.Conn) {
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read msg error" + err.Error())
			return
		}
		if messageType == websocket.PingMessage {
			fmt.Println("this is a  pingMessage" + err.Error())
		}
		ws.writeMsg(conn, messageType, msg)
	}
}

func (ws *WServer) writeMsg(conn *websocket.Conn, msgType int, msg []byte) error {
	return conn.WriteMessage(msgType, msg)
}
