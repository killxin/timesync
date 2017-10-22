package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"

	"github.com/killxin/timesync/proto"
)

var serverIP = "localhost"
var serverPort = ":9527"
var pwd = "qscvgyuk./'][;.,kuygvcdw"

func main() {
	client, err := rpc.DialHTTP("tcp", serverIP+serverPort)
	if err != nil {
		log.Fatal("dialing", err)
	}
	args := &proto.Args{PWD: pwd}
	reply := &proto.Reply{}
	for {
		t1 := time.Now().UnixNano()
		err = client.Call("Clock.Sync", args, &reply)
		t2 := time.Now().UnixNano()
		if err != nil {
			log.Fatal("sync error: ", err)
		}
		delay := (t2 - t1 - (reply.T1 - reply.T2)) / 2
		t := time.Unix(0 /*sec*/, reply.T2+delay /*nsec*/)
		fmt.Printf("%s delay=%dns\n", t, delay)
		// clock sync pre 10 seconds
		display(t, 10)
	}
}

func display(t time.Time, count int) {
	ticker := time.NewTicker(time.Second)
	i := 0
	for range ticker.C {
		t = t.Add(time.Second)
		fmt.Printf("%s\n", t)
		i++
		if i == count {
			ticker.Stop()
			return
		}
	}
}
