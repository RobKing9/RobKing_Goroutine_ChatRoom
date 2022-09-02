package main

import (
	"log"
)

//接收数据
func (s *Server) recData(cli *OnlineClient) {
	for {
		n, err := cli.Conn.Read(cli.Buf)
		if err != nil {
			cli.Conn.Close()
		}
		cli.RecData <- string(cli.Buf[:n])
	}
}

//发送数据
func (s *Server) sentData(cli *OnlineClient) {
	//不断地从管道中拿出数据 发送给客户端
	for {
		data := <-cli.SentData
		_, err := cli.Conn.Write([]byte(data))
		if err != nil {
			cli.Conn.Close()
			log.Println("conn.Write failed:", err.Error())
		}
	}
}

//获取当前在线用户名
func (s *Server) onlineList(name string) string {
	if len(s.OnlineClients) <= 1 {
		return "当前没有人在线!"
	} else {
		onlineList := "当前在线的用户有:"
		for onlineName := range s.OnlineClients {
			//除去 自己
			if onlineName != name {
				onlineList = onlineList + onlineName + ","
			}
		}
		return onlineList + "快去找他们聊天吧!"
	}
}

//将当前用户的信息 放入管道
func (s *Server) recMessage(cli OnlineClient) {
	message := <-cli.RecData
	if len(message) != 0 {
		s.Message <- message
	}
}

//给当前用户发送消息
//func (s *Server) sentMessage(cli OnlineClient) {
//	for msg := range cli.CliChanMessage {
//		cli.Conn.Write([]byte(msg))
//	}
//}
