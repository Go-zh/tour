// +build no-build OMIT

package main

import "github.com/Go-zh/tour/reader"

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法

func main() {
	reader.Validate(MyReader{})
}
