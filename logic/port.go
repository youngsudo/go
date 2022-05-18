package logic

import (
	"fmt"
	"net"
)

func FindPort(port int) int {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", port)) // 解析TCP地址
	if err != nil {
		return FindPort(port + 1)
	}
	l, err := net.ListenTCP("tcp", addr) // 监听TCP地址
	if err != nil {
		return FindPort(port + 1)
	}
	l.Close()
	return port
}
