package main

import "log"

//接收数据
func (s *Server) recData(c *OnlineClient) {
	for {
		n, err := c.Conn.Read(c.buf)
		if err != nil {
			log.Println("conn.Read failed:", err.Error())
		}
		c.RecData <- string(c.buf[:n])
	}
}

//发送数据
func (s *Server) sentData(c *OnlineClient) {
	//不断地从管道中拿出数据 发送给客户端
	for {
		data := <-c.SentData
		_, err := c.Conn.Write([]byte(data))
		if err != nil {
			log.Println("conn.Write failed:", err.Error())
		}
	}
}

//获取当前在线用户名
func (s *Server) onlineList(name string) string {
	if len(s.OnlineClients) == 0 {
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

//接收当前用户 的消息并且广播出去
func (s *Server) recMessage(cli OnlineClient) {
	message := <-cli.RecData
	if len(message) != 0 {
		s.Message <- message
	}
}
