package main

import (
	"log"
)

//发送数据
func (c *Client) sentData() {
	//不断地从管道中拿出数据 发送给服务器
	for {
		data := <-c.SentData
		_, err := c.Conn.Write([]byte(data))
		if err != nil {
			log.Println("conn.Write failed:", err.Error())
		}
	}
}

//接收数据
func (c *Client) recData() {
	for {
		n, err := c.Conn.Read(c.buf)
		if err != nil {
			c.Conn.Close()
			log.Println("conn.Read failed:", err.Error())
			return
		}
		c.RecData <- string(c.buf[:n])
	}
}
