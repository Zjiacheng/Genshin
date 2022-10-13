package game

import (
	"YuanShen_Server/src/csvs"
	"fmt"
	"math/rand"
	"time"
)

type Server struct {
	Num int
}

var server *Server

func GetServer() *Server {
	if server == nil {
		server = new(Server)
	}
	return server
}

func (self *Server) Start() {
	csvs.CheckLoadCsv()
	rand.Seed(time.Now().Unix())
	fmt.Printf("-----------------Test----------------\n")
	player := NewTestPlayer()
	go player.Run()
	for {
		if self.Num == 0 {
			break
		}
	}
	fmt.Println("服务器正常关闭！")
}
func (self *Server) Close() {

}

func (self *Server) AddGo() {
	self.Num++
}

func (self *Server) GoDone() {
	self.Num--
}
