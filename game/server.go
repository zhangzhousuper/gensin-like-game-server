package game

import (
	"fmt"
	"gensin-server/csvs"
	"math/rand"
	"os"
	"os/signal"
	"regexp"
	"sync"
	"syscall"
	"time"
)

type Server struct {
	Wait        sync.WaitGroup
	BanWordBase []string
	Lock        *sync.RWMutex
}

var server *Server

func GetServer() *Server {
	if server == nil {
		server = new(Server)
		server.Lock = new(sync.RWMutex)
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
	go self.SignalHandle()

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

func (self *Server) IsBanWord(txt string) bool {
	self.Lock.RLock()
	defer self.Lock.RUnlock()
	for _, v := range self.BanWordBase {
		match, _ := regexp.MatchString(v, txt)
		if match {
			fmt.Println("发现违禁词:", v)
		}
		if match {
			return match
		}
	}
	return false
}

func (self *Server) UpdateBanWord(banWord []string) {
	self.Lock.Lock()
	defer self.Lock.Unlock()
	self.BanWordBase = banWord
}

func (self *Server) SignalHandle() {
	channelSignal := make(chan os.Signal)
	signal.Notify(channelSignal, syscall.SIGINT)

	for {
		select {
		case <-channelSignal:
			self.Close()
		}
	}
}
