package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Client struct {
	Conn     net.Conn    //连接对象
	Name     string      //用户名
	Password string      //密码
	RecData  chan string //接收信息的管道
	buf      []byte      //接收数据缓存
	SentData chan string //发送信息的管道
}

func (c *Client) MakeClient() {
	//初始化客户端信息
	//初始化管道
	c.RecData = make(chan string)
	c.buf = make([]byte, 1024)
	c.SentData = make(chan string)
	//连接服务器
	conn, err := net.Dial("tcp", ":9999")
	c.Conn = conn
	if err != nil {
		log.Println("net.dail failed:", err.Error())
	}
}

func main() {
	//初始化客户端
	client := &Client{}
	client.MakeClient()
	time.Sleep(2 * time.Second)
	log.Println("连接服务器成功!")
	//启动数据接受和发送 协程
	go client.recData()
	go client.sentData()
	//登录注册
	log.Print("请输入将要进行的操作：1、登录  2、注册")
	for {
		var input string
		fmt.Scanln(&input)
		if input == "1" {
			//登录
			client.clientLogin()
			break
		} else if input == "2" {
			//注册
			client.clientRegister()
			break
		} else {
			log.Println("输入有误!请重新输入!!!")
		}
	}

}
