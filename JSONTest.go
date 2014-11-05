package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func JsonMarshal() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

func JsonUnmarshal() {
	var s Serverslice
	str := `{"servers":
  		  [{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},
  		   {"serverName":"Beijing_VPN","serverIP":"127.0.0.1"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}

func main() {
	fmt.Println("-------JsonUnmarshal---------")
	JsonUnmarshal()
	fmt.Println("--------JsonMarshal----------")
	JsonMarshal()
}
