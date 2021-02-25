package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"net"
	"regexp"
	"strconv"
)

func main() {

	var (
		mac  string
		host string
		port int
	)

	broadcast := "255.255.255.255"

	flag.StringVar(&mac, "mac", "", "MAC地址\n例如：\n00:00:00:00:00:00\n00-00-00-00-00-00\nAABBCCDDEEFF")
	flag.StringVar(&host, "host", broadcast, "主机地址")
	flag.IntVar(&port, "port", 9, "端口号")

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		args = make([]string, 3) // make一个空数组，防止下面的代码报索引错误
	}

	if mac == "" && args[0] != "" {
		mac = args[0]
	}
	if mac == "" && args[0] == "" {
		fmt.Print("MAC地址不能为空")
		return
	}

	ret, _ := regexp.MatchString("^([A-Fa-f0-9]{2}[-,:]){5}[A-Fa-f0-9]{2}$", mac)
	if !ret && len(mac) != 12 {
		fmt.Print("MAC地址格式错误")
		return
	}

	if (host == "" || host == broadcast) && args[1] != "" {
		host = args[1]
	}
	if host == "" && args[1] == "" {
		fmt.Print("主机地址不能为空")
		return
	}

	if port == 0 && args[2] != "" {
		port, _ = strconv.Atoi(args[2])
	}

	buffer := bytes.Buffer{}
	for i := 0; i < 6; i++ {
		buffer.WriteByte(0xFF)
	}
	reg, _ := regexp.Compile("[-,:]")
	mac = reg.ReplaceAllString(mac, "")
	macBytes, _ := hex.DecodeString(mac)
	for i := 0; i < 16; i++ {
		buffer.Write(macBytes)
	}
	addr := net.UDPAddr{
		IP:   net.ParseIP(host),
		Port: port,
	}
	conn, err := net.DialUDP("udp", nil, &addr)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer conn.Close()

	sLen, err := conn.Write(buffer.Bytes())
	if err != nil {
		fmt.Print(err)
		return
	}
	if sLen > 0 {
		fmt.Print("WOL数据包发送成功")
	}
}
