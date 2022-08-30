package main

import (
	"fmt"
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
	var input string
	fmt.Print("输入消息以发送:")
	for {
		fmt.Scanln(&input)
		//退出聊天室
		if input == "quit" {
			c.SentData <- c.Name + ":" + input
			input = ""
			log.Println("欢迎下次使用!")
			os.Exit(0)
		}
		//发送消息
		if len(input) != 0 {
			c.SentData <- c.Name + ":" + input
			input = ""
		}
	}
}
