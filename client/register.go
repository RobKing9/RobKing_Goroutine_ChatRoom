package main

import (
	"fmt"
	"log"
)

//用户注册
func (c *Client) clientRegister() {
	for {
		//用户名，密码
		var name, psw string
		fmt.Print("请输入用户名:")
		fmt.Scanln(&name)
		fmt.Print("请输入密码:")
		fmt.Scanln(&psw)
		//将用户名和密码发给服务器
		c.SentData <- fmt.Sprintf("用户名-%s-密码-%s-请求注册!", name, psw)
		//服务器同意注册
		info := <-c.RecData
		//注册成功
		if info == "success" {
			log.Println("注册成功!请登录!")
			break
		} else {
			log.Printf("注册失败:%s, 请重新输入!", info)
		}
	}
	c.clientLogin()
}
