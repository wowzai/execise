package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName   xml.Name `xml:"servers"`
	Version   string `xml:"version,attr"`
	Svs       []server `xml:"server"`
	Comment   string   `xml:",comment"`
	Description string `xml:",innerxml"`
}

type server struct {
	XMLName   xml.Name `xml:"server"`
	ServerName string  `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

type Servers struct {
	XMLName  xml.Name `xml:"servers"`
	Version  string   `xml:"version,attr"`
	Svs      []serv   `xml:"server"`
}

type serv struct {
	ServerName string `xml:"ServerName"`
	ServerIP   string `xml:"ServerIP"`
}

func main() {
	servers := &Servers{Version:"1"}
	servers.Svs = append(servers.Svs,serv{"Shanghai_VPN","127.0.0.1"})
	servers.Svs = append(servers.Svs,serv{"Beijing_VPN","127.0.0.1"})
	output,err := xml.MarshalIndent(servers," ","  ")
	if err != nil {
		fmt.Printf("error:%v\n",err)
	}
	os.Stdout.Write([]byte(xml.Header))

	os.Stdout.Write(output)

	fmt.Println()

	file,err := os.Open("server.xml")
	if err != nil {
		fmt.Printf("error:%v",err)
		return
	}
	defer file.Close()
	data,err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error:%v",err)
		return
	}
    v := Recurlyservers{}
    err = xml.Unmarshal(data,&v)
    if err != nil {
    	fmt.Printf("error:%v",err)
    	return
    }

    fmt.Printf("comment:%s\n",v.Comment)

    fmt.Println(v)
}