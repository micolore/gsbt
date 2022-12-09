package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	download()
}

func urlEncoder() {
	u, _ := url.Parse("http://www.example.com/hello world")
	fmt.Println(u.Path)
	fmt.Println(u.Path)
	u.Path = url.PathEscape(u.Path)
	fmt.Println(u)
}

func download() {
	// 请求文件，并获取响应
	resp, err := http.Get("http://.com/temp/aga in.txt")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 创建文件
	file, err := os.Create("examp le.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 将响应写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
}
