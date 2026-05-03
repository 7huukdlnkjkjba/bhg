package main

import (
	"fmt"
	"log"
	"os"
)

// FooReader定义一个从stdin读取的io.Reader。
type FooReader struct{}

// Read从stdin读取数据。
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

// FooWriter定义一个写入Stdout的io.Writer。
type FooWriter struct{}

// Write写入数据到Stdout。
func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func main() {
	// 实例化reader和writer。
	var (
		reader FooReader
		writer FooWriter
	)

	// 创建缓冲区来保存输入/输出。
	input := make([]byte, 4096)

	// 使用reader读取输入。
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	// 使用writer写入输出。
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)
}
