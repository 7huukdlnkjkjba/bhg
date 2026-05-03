package main

import (
	"fmt"
	"io"
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

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}
