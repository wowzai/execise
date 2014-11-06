package main

import (
  "fmt"
  "net"
  "os"
  "time"
)

func main() {
	service := ":1200"
	udpAddr,err := net.ResolveUDPAddr("udp4",service)
	checkErr(err)
	conn,err := net.ListenUDP("udp",udpAddr)
	checkErr(err)
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_,addr,err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime),addr)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr,"Fatal error: %s",err.Error())
		os.Exit(1)
	}
}