package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	databases "moppo.com/gsbt/mysql"
	"os"
	"strings"
)

type ReqInfo struct {
	Url     string
	FullUrl string
	Time    string
}

func SaveReqInfo(reqInfo ReqInfo) {
	databases.DB.Save(&reqInfo)
}
func main() {
	getAgainData()
}

func getAgainData() {
	filePath := "/Users/kubrick/Documents/wscrm-bus-api"
	inputFile, inputError := os.Open(filePath)
	if inputError != nil {
		fmt.Printf("打开文件时出错", inputError.Error())
	}
	inputReader := bufio.NewReader(inputFile)
	pageSize := 50
	i := 0
	sqlValues := make([]ReqInfo, 10)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		}
		lineContent := strings.Split(inputString, " ")
		urls := strings.Split(lineContent[4], "?")
		time := lineContent[1]
		fullUrl := lineContent[4]
		var reqi ReqInfo

		realTime := strings.ReplaceAll(time, "[", "")
		reqi.Url=urls[0]
		reqi.FullUrl=fullUrl
		reqi.Time=realTime
		sqlValues = append(sqlValues, reqi)
		i++
		if i == pageSize {
			batchInsert(sqlValues)
			i = 0
		}
	}
	defer inputFile.Close()
}
func batchInsert(arrs []ReqInfo) {
	basicSql := "insert into t_req (url,time,full_url) values"
	var buffer bytes.Buffer

	if _, err := buffer.WriteString(basicSql); err != nil {
		return
	}
	for i, e := range arrs {
		if e.FullUrl == ""{
			continue
		}
		if i == len(arrs)-1 {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%s');", e.Url, e.FullUrl, e.Time))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%s'),", e.Url, e.FullUrl, e.Time))
		}
	}
	tx := databases.DB.Exec(buffer.String())
	if tx.Error != nil {
		panic(tx.Error)
	}
}
func arrayToString(arr []string) string {
	var result string
	for _, i := range arr {
		result += i
	}
	return result
}

func writeFile(content string) {
	outputFile, outputError := os.OpenFile("/Users/kubrick/Documents/20210926-req.sql", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString(content)
	outputWriter.Flush()
}
