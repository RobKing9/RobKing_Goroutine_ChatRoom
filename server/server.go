package main

import (
	"github.com/jinzhu/gorm"
	"log"
	"net"
)

type Server struct {
	Db      *gorm.DB     //连接对象
	listen  net.Listener //监听
	Message chan string  //广播消息的管道

	OnlineClients map[string]OnlineClient //所有用户的信息
}

// OnlineClient 在线用户
type OnlineClient struct {
	Name     string      //名称
	Conn     net.Conn    //连接对象
	RecData  chan string //接收信息的管道
	SentData chan string //给单独客户端发送信息的管道
	Buf      []byte
}

func (s *Server) MakeServer() {
	//初始化 服务器
	//初始化管道
	s.Message = make(chan string, 1024)
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
		conn, err := server.listen.Accept()
		if err != nil {
			log.Println("l.Accept failed:", err.Error())
		}
		log.Printf("用户%s连接服务器成功！", conn.RemoteAddr().String())
		//处理用户连接
		go server.handleConn(conn)
	}
}

//处理每一个连接
func (s *Server) handleConn(conn net.Conn) {
	//初始化 连接对象
	cli := OnlineClient{}
	cli.Conn = conn
	cli.SentData = make(chan string)
	cli.RecData = make(chan string)
	cli.Buf = make([]byte, 1024)
	////用户的数据接收和发送协程
	go s.recData(&cli)
	go s.sentData(&cli)
	////处理用户登录注册
	go s.clientAuth(cli)
}
