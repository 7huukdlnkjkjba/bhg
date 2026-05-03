package main

import (
	"io"
	"log"
	"net"
)

// echo是一个简单的回显接收数据的处理函数。
func echo(conn net.Conn) {
	defer conn.Close()

	// 创建一个缓冲区来存储接收的数据。
	b := make([]byte, 512)
	for {
		// 通过conn.Read将数据接收至缓冲区。
		size, err := conn.Read(b[0:])
		if err != nil && err != io.EOF {
			log.Println("Unexpected error")
			break
		}

		if err == io.EOF && size == 0 {
			log.Println("Client disconnected")
			break
		}

		log.Printf("Received %d bytes: %s", size, string(b))

		// 通过conn.Write发送数据。
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	// 绑定到所有接口的TCP端口20080。
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		// 等待连接。在连接建立时创建net.Conn。
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// 处理连接。使用goroutine实现并发。
		go echo(conn)
	}
}
