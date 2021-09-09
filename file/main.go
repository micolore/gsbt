package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//read dir  java file out office
	//content := readSingleFile("/Users/kubrick/Downloads/tjava/mengxuegu-blog/mengxuegu-blog-api/src/main/java/com/mengxuegu/blog/feign/req/UserInfoREQ.java")
	//content := readSingleFile(path)
	//fmt.Println(content)
	//writeFile(content)

	path := "/Users/kubrick/Documents/moppo/code/peoject/scrm-2.0/whatsapp-scrm-server/scrm-admin-server-modules/"
	GetAllFile(path)

	path = "/Users/kubrick/Documents/moppo/code/peoject/scrm-2.0/whatsapp-scrm-server/scrm-biz-modules/"
	GetAllFile(path)

	path = "/Users/kubrick/Documents/moppo/code/peoject/scrm-2.0/whatsapp-scrm-server/scrm-common-modules/"
	GetAllFile(path)

	path = "/Users/kubrick/Documents/moppo/code/peoject/scrm-2.0/whatsapp-scrm-server/scrm-single-server-modules/"
	GetAllFile(path)

}

func writeDoc() {

}

// GetAllFile 获取目录下面所有的文件（递归获取）
func GetAllFile(path string) {
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, fi := range rd {
		fn := fi.Name()
		if fi.IsDir() {
			GetAllFile(path + fn + "//")
		} else {
			if strings.Contains(fn, ".java") {
				fullName := path + fn
				fmt.Println(fullName)
				content := readSingleFile(fullName)
				writeFile(content)
			}
		}
	}
}

func writeFile(content string) {
	outputFile, outputError := os.OpenFile("/Users/kubrick/Documents/20210909-code.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString(content)
	outputWriter.Flush()
}

func readSingleFile(filePath string) string {
	inputFile, inputError := os.Open(filePath)
	if inputError != nil {
		fmt.Printf("打开文件时出错", inputError.Error())
		return ""
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	fullFileString := ""
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return fullFileString
		}
		//if inputString == "\r\n" {
		//	continue
		//}
		inputString = strings.Trim(inputString, "\r\n")
		fullFileString += inputString + "\n"
	}
	return fullFileString
}
