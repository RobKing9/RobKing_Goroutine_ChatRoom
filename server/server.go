package main

import (
	"github.com/jinzhu/gorm"
	"log"
	"net"
)

type Server struct {
	Db            *gorm.DB                //连接对象
	listen        net.Listener            //监听
	Message       chan string             //广播消息的管道
	OnlineClients map[string]OnlineClient //所有用户的信息
}

// OnlineClient 在线用户
type OnlineClient struct {
	Conn     net.Conn    //连接对象
	RecData  chan string //接收信息的管道
	buf      []byte
	SentData chan string //给单独客户端发送信息的管道
}

func (s *Server) MakeServer() {
	//初始化 服务器
	//初始化管道
	s.Message = make(chan string)
	s.OnlineClients = make(map[string]OnlineClient)
	//连接mysql
	err := s.initMysql()
	if err != nil {
		log.Println("连接数据库失败!")
	}
	s.Db.AutoMigrate(&ClientInfo{})
	log.Println("连接数据库成功!")
	//监听端口
	l, err := net.Listen("tcp", ":9999")
	s.listen = l
	log.Println("服务器启动成功!正在监听端口9999.....")
	if err != nil {
		log.Println("net.Listen failed:", err.Error())
	}
}

func main() {
	server := &Server{}
	server.MakeServer()
	//广播消息 转发消息
	go server.broadcast()
	//持续等待连接
	for {
		client := OnlineClient{}

		conn, err := server.listen.Accept()
		client.Conn = conn
		if err != nil {
			log.Println("l.Accept failed:", err.Error())
		}
		//初始化连接的客户端
		client.RecData = make(chan string)
		client.SentData = make(chan string)
		client.buf = make([]byte, 1024)
		log.Printf("用户%s连接服务器成功！", conn.RemoteAddr().String())
		//该用户的数据接收和发送协程
		go server.recData(&client)
		go server.sentData(&client)
		//处理用户登录注册
		go server.clientAuth(client)
	}
}
