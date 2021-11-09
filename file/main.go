package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GenCode() {
	//read dir  java file out office
	//content := readSingleFile("/Users/kubrick/Downloads/tjava/mengxuegu-blog/mengxuegu-blog-api/src/main/java/com/mengxuegu/blog/feign/req/UserInfoREQ.java")
	//content := readSingleFile(path)
	//fmt.Println(content)
	//writeFile(content)

	// 1)
	// path := "/Users/kubrick/Documents/moppo/code/peoject/scrm-2.0/whatsapp-scrm-server/scrm-admin-server-modules/"
	// path = "/Users/kubrick/Documents/moppo/code/peoject/scrm-2.0/whatsapp-scrm-server/scrm-single-server-modules/"
	// path = "/Users/kubrick/Documents/moppo/code/peoject/scrm-2.0/whatsapp-scrm-server/scrm-biz-modules/"
	// path = "/Users/kubrick/Documents/moppo/code/peoject/scrm-2.0/whatsapp-scrm-server/scrm-common-modules/"

	// 2)
	path := "/Users/kubrick/Documents/moppo/code/peoject/marketing-treasure-server/marketing-treasure-server/marketing-api/"
	GetAllFile(path)
	path = "/Users/kubrick/Documents/moppo/code/peoject/marketing-treasure-server/marketing-common-modules/"
	GetAllFile(path)
	path = "/Users/kubrick/Documents/moppo/code/peoject/marketing-treasure-server/marketing-treasure-server/marketing-manager/"
	GetAllFile(path)
}

func main() {
	HandleNginxLog()
}

func HandleNginxLog() {
	filePath := "/Users/kubrick/Downloads/2021-09-25.log"
	inputFile, inputError := os.Open(filePath)
	if inputError != nil {
		fmt.Printf("打开文件时出错", inputError.Error())
	}
	inputReader := bufio.NewReader(inputFile)
	var countMap map[string]ReqInfo
	countMap = make(map[string]ReqInfo)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		}
		lineContent := strings.Split(inputString, " ")
		urls := strings.Split(lineContent[4], "?")
		time := lineContent[1]
		fullUrl := lineContent[4]

		url := urls[0]
		value, ok := countMap[url]
		if ok {
			value.Count = value.Count + 1
			countMap[url] = value
		} else {
			fullUrl=""
			value := ReqInfo{url, 0, fullUrl, time}
			countMap[url] = value
		}
	}
	var keys []string
	for k := range countMap {
		keys = append(keys, k)
	}
	var reqInfos = make([]ReqInfo, 1)
	i := 0
	for _, k := range keys {
		reqInfos = append(reqInfos, countMap[k])
		i++
	}
	sort.Sort(ReqInfoSliceDecrement(reqInfos))
	fmt.Println(reqInfos)
	defer inputFile.Close()
}

type ReqInfo struct {
	Url     string
	Count   int
	FullUrl string
	Time    string
}
type ReqInfoSliceDecrement []ReqInfo

func (s ReqInfoSliceDecrement) Len() int { return len(s) }

func (s ReqInfoSliceDecrement) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s ReqInfoSliceDecrement) Less(i, j int) bool { return s[i].Count > s[j].Count }

// GetAllFile 获取目录下面所有的文件（递归获取）
func GetAllFile(path string) {
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, fi := range rd {
		fn := fi.Name()
		if fi.IsDir() {
			GetAllFile(path + fn + "/")
		} else {
			if strings.Contains(fn, ".java") {
				fullName := path + fn
				content := readSingleLineFile(fullName)
				fmt.Println(fullName)
				writeFile(content)
			}
		}
	}
}
func writeFile(content string) {
	outputFile, outputError := os.OpenFile("/Users/kubrick/Documents/20210926-code.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString(content)
	outputWriter.Flush()
}

// 优化版本
func writeFileV2(content string) {
	fileName := "/Users/kubrick/Documents/20210926-code.txt"
	var outputFile *os.File

	outputFile, err := os.OpenFile(fileName, os.O_RDONLY|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Open file err =", err)
		return
	}
	if outputFile == nil {
		var outputError error
		outputFile, outputError = os.OpenFile(fileName, os.O_WRONLY, 0644)
		if outputError != nil {
			fmt.Printf("An error occurred with file opening or creation\n")
			return
		}
	}
	defer outputFile.Close()
	_, err = io.WriteString(outputFile, content)
	if err != nil {
		panic(err)
	}
}

func readSingleLineFile(filePath string) string {
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

//  切割大日志文件 50m一个
func SpitFile(filePath string) {

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileSize int64 = fileInfo.Size()

	const fileChunk = 50 * (1 << 20) // 1 MB

	totalPartsNum := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	fmt.Printf("Splitting to %d pieces.\n", totalPartsNum)

	for i := uint64(1); i < totalPartsNum; i++ {

		partSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))
		partBuffer := make([]byte, partSize)

		file.Read(partBuffer)

		fileName := "slow-log_" + strconv.FormatUint(i, 10)
		_, err := os.Create(fileName)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ioutil.WriteFile(fileName, partBuffer, os.ModeAppend)
		fmt.Println("Split to : ", fileName)
	}
}
