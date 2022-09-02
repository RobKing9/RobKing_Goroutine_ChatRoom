package main

import (
	"log"
	"strings"
)

//广播消息 转发消息
func (s *Server) broadcast() {
	for {
		message := <-s.Message
		log.Println(message)
		//处理用户的信息
		//用户发送 “quit”
		m := strings.Split(message, ":")
		var name, msg string
		if len(m) == 2 {
			name, msg = m[0], m[1]
		} else {
			msg = m[0]
		}
		if msg == "quit" {
			//在线用户缓存中 去除
			delete(s.OnlineClients, name)
			//通知其他用户 该用户下线
			s.Message <- "用户:" + name + "已经下线!" + s.onlineList(name)
		} else if msg[:2] == "to" && len(msg) > 4 { //处理私聊信息	to name content
			//前面分离出来的事 指定客户端的名称 后面的是消息内容
			s.OnlineClients[strings.Split(msg, " ")[1]].SentData <- name + " 私聊你说: " + strings.Split(msg, " ")[2]
		} else {
			//给每一个在线用户发送消息	广播的功能
			for _, cli := range s.OnlineClients {
				cli.SentData <- message
			}
		}
	}
}
