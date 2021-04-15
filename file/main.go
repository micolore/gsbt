package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// read dir  java file out office
	//content := readSingleFile("/Users/kubrick/Downloads/tjava/mengxuegu-blog/mengxuegu-blog-api/src/main/java/com/mengxuegu/blog/feign/req/UserInfoREQ.java")
	//writeFile(content)
	writeDoc()
}

func writeDoc() {

}

func writeFile(content string) {
	outputFile, outputError := os.OpenFile("/Users/kubrick/Downloads/2.doc", os.O_WRONLY|os.O_CREATE, 0666)
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
