package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type ClientInfo struct {
	gorm.Model
	Name string
	Psw  string
}

//数据库相关的操作
func (s *Server) initMysql() (err error) {
	dsn := "root:@tcp/ChatRoom?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Println("gorm.Open failed:", err.Error())
	}
	s.Db = db
	return s.Db.DB().Ping()
}

//添加客户端
func (s *Server) addClient(cli *ClientInfo) (err error) {
	err = s.Db.Create(cli).Error
	return
}

//通过name查找
func (s *Server) getClientByName(name string) (*ClientInfo, error) {
	var cli ClientInfo
	if err := s.Db.Where("name = ?", name).First(&cli).Error; err != nil {
		return nil, err
	}
	return &cli, nil
}

//判断客户端是否存在
func (s *Server) isClientExist(name string) bool {
	var cli ClientInfo
	s.Db.Where("name = ?", name).First(&cli)
	if cli.Name != "" {
		return true
	}
	return false
}
