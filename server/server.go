package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strings"
	"time"
	"flag"

	"github.com/killxin/timesync/proto"
)

var pwd = "qscvgyuk./'][;.,kuygvcdw"

// ClockImpl implements Clock Interface
type ClockImpl int64

// Sync implements Clock Interface
func (t *ClockImpl) Sync(args *proto.Args, reply *proto.Reply) error {
	reply.T1 = time.Now().UnixNano()
	if strings.Compare(args.PWD, pwd) != 0 {
		return errors.New("Authorization Failed")
	}
	time.Sleep(3 * time.Millisecond)
	reply.T2 = time.Now().UnixNano()
	return nil
}

func main() {
	serverPort := flag.String("port", ":9527", "http listen port")
	flag.Parse()
	var clock proto.Clock
	clock = new(ClockImpl)
	rpc.RegisterName("Clock", clock)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", *serverPort)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		fmt.Println(t)
	}
	// for timer := time.NewTimer(time.Second);; {
	// 	select {
	// 	case <-timer.C:
	// 		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// 	}
	// 	timer = time.NewTimer(time.Second)
	// }
}
