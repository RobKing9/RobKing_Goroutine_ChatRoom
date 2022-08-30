package main

import (
	"fmt"
	"log"
)

//用户登录
func (c *Client) clientLogin() {
	for {
		//用户名，密码
		var name, psw string
		fmt.Print("请输入用户名:")
		fmt.Scanln(&name)
		fmt.Print("请输入密码:")
		fmt.Scanln(&psw)
		//发送给服务器进行验证
		c.SentData <- fmt.Sprintf("用户名-%s-密码-%s-请求登录!", name, psw)
		info := <-c.RecData
		//验证成功
		if info == "success" {
			fmt.Println("登录成功!欢迎进入聊天室!\n输入“quit”退出聊天室\n输入“to 对象 ”给指定对象发送消息")
			c.Name = name
			c.Password = psw
			//接收 并显示消息协程
			go c.recMessage()
			break
		} else {
			log.Println("登录失败:", info)
		}
	}
	//聊天内容输入协程
	c.sentMessage()
}
