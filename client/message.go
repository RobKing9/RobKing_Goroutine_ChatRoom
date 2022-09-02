package main

import (
	"bufio"
	"log"
	"os"
)

//处理连接后的消息

//接收消息 显示在界面
func (c *Client) recMessage() {
	for {
		message := <-c.RecData
		log.Println(message)
	}
}

//发送消息
func (c *Client) sentMessage() {
	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		//退出聊天室
		if input == "quit" {
			c.SentData <- c.Name + ":" + input
			input = ""
			log.Println("欢迎下次使用!")
			return
		}
		//发送消息
		if len(input) != 0 {
			c.SentData <- c.Name + ":" + input
			input = ""
		}
	}
}
