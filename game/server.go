package game

import (
	"fmt"
	"gensin-server/csvs"
	"math/rand"
	"sync"
	"time"
)

type Server struct {
	Num  int
	lock sync.Mutex
	Wait sync.WaitGroup
}

var server *Server

func GetServer() *Server {
	if server == nil {
		server = new(Server)
	}
	return server
}

func (self *Server) Start() {
	//**********************************************************
	// 加载配置
	rand.Seed(time.Now().Unix())
	csvs.CheckLoadCsv()
	go GetManageBanWord().Run()

	//fmt.Printf("数据测试----start\n")
	playerTest := NewTestPlayer()
	go playerTest.Run()

	self.Wait.Wait()
	fmt.Println("服务器关闭")
}

func (self *Server) Close() {
	GetManageBanWord().Close()
}

func (self *Server) AddGo() {
	self.Wait.Add(1)
}

func (self *Server) GoDone() {
	self.Wait.Done()
}
