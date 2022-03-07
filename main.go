/*
 * @Description: Add file description
 * @Author: lilihx@github.com
 * @Date: 2022-03-07 13:57:21
 * @LastEditTime: 2022-03-07 17:05:48
 * @LastEditors: lilihx@github.com
 */
package main

import "github.com/lilihx/chatRoom/wss"

func main() {
	ser := wss.WServer{}
	ser.InitWebSocket()
}
