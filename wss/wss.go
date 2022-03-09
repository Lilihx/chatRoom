/*
 * @Description: Add file description
 * @Author: lilihx@github.com
 * @Date: 2022-03-04 16:43:48
 * @LastEditTime: 2022-03-08 15:28:04
 * @LastEditors: lilihx@github.com
 */
package wss

import (
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/lilihx/chatRoom/common/config"
	"github.com/lilihx/chatRoom/common/discover"
	"github.com/lilihx/chatRoom/common/log"
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

	consulClient, err := discover.NewKitDiscoverClient(config.Config.Consul.Host, config.Config.Consul.Port)
	if err != nil {
		log.Error("Init consul client err:" + err.Error())
		return err
	}
	consulClient.Register("wss", "123456", "/health", config.Config.WServer.Host, config.Config.WServer.Port, nil)
	defer consulClient.DeRegister("123456")

	url := config.Config.WServer.Host + ":" + strconv.Itoa(config.Config.WServer.Port)
	err = http.ListenAndServe(url, nil)

	if err != nil {
		log.Error("Init websocket err:" + err.Error())
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
		log.Error("init websocket err")
		return
	}
	go ws.readMsg(conn)
}

func (ws *WServer) readMsg(conn *websocket.Conn) {
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Error("read msg error" + err.Error())
			return
		}
		if messageType == websocket.PingMessage {
			log.Info("this is a  pingMessage" + err.Error())
		}
		ws.writeMsg(conn, messageType, msg)
	}
}

func (ws *WServer) writeMsg(conn *websocket.Conn, msgType int, msg []byte) error {
	return conn.WriteMessage(msgType, msg)
}
