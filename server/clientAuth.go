package main

import (
	"log"
	"strings"
)

//处理用户登录和注册
func (s *Server) clientAuth(c OnlineClient) {
	for {
		//等待用户登录注册
		info := strings.Split(<-c.RecData, "-")
		name, psw, flag := info[1], info[3], info[4]
		if flag == "请求注册!" {
			//处理注册
			//log.Printf("用户名:%s, 密码:%s 请求注册!", name, psw)
			_, info := s.handleRegister(name, psw)
			c.SentData <- info
		} else {
			//处理登录
			//log.Printf("用户名:%s, 密码:%s 请求登录!", name, psw)
			status, info := s.handleLogin(name, psw)
			c.SentData <- info
			//登录成功
			if status == true {
				s.Message <- "欢迎用户:" + name + "进入聊天室!"
				//更新在线用户
				s.OnlineClients[name] = c
				//给这个用户 发送当前在线用户
				c.SentData <- s.onlineList(name)
				//接收该用户发的消息
				go s.recMessage(c)
				break
			}
		}
	}
}

//处理用户注册
//逻辑：通过name从数据库中 判断用户名是否存在，如果不存在保存到数据库即可
func (s *Server) handleRegister(name, psw string) (status bool, info string) {
	//首先通过name查找是否存在
	if s.isClientExist(name) {
		return false, "用户名已经存在!"
	} else {
		//如果不存在 保存到数据库中
		var client ClientInfo
		client.Name = name
		client.Psw = psw
		err := s.addClient(&client)
		if err != nil {
			log.Println("addClient failed:", err.Error())
		}
		return true, "success"
	}
}

//处理用户登录
//逻辑：和通过name在数据库中查找 进行密码比对
func (s *Server) handleLogin(name, psw string) (status bool, info string) {
	//和数据库中的账号密码进行对比
	cli, err := s.getClientByName(name)
	if err != nil {
		log.Println("getClientByName failed:", err.Error())
	}
	//密码比对失败
	if cli.Psw != psw {
		return false, "密码输入错误!"
	} else {
		return true, "success"
	}
}
