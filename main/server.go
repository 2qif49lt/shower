package main

import (
	"fmt"
	"github.com/2qif49lt/shower/httpsrv"
	"github.com/2qif49lt/shower/proxys"
	tcpsrv "github.com/2qif49lt/shower/tcp/server"
	"github.com/2qif49lt/shower/udpsrv"
)

func main() {
	udpSrv := udpsrv.New()
	ps := proxys.New()
	ps.Udpsender = udpSrv
	udpSrv.SetHander(&udpsrv.MsgHandler{ps})
	go httpsrv.RunOn(":9090", ps, ":7898")
	tcpSrv := tcpsrv.New(ps)
	go tcpSrv.RunOn(":7898")
	err := udpSrv.RunOn(":8898")
	fmt.Println(err)

}
