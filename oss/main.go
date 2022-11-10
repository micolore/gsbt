package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// oss 上传文件
// 1、命令参数-指定目录下面文件，上传到指定目录下面
// 2、新增配置文件配置oss配置信息
// 3、指定文件上传到指定目录下面
func main() {
	files, err := GetAllFiles("/Users/kubrick/Documents/项目文档")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(files); i++ {
		file := files[i]
		fmt.Println(files[i])
		// linux
		// uploadFile(strings.TrimPrefix(file, "/"), file)

		// windows
		//获取文件名带后缀
		filenameWithSuffix := path.Base(file)
		uploadFile("kubrick/20221110/"+filenameWithSuffix, file)

	}
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

const endpoint = "oss-cn-shanghai.aliyuncs.com"
const accessKeyId = ""
const accessKeySecret = ""
const bucketName = "one-return"

func uploadFile(objectName, localFileName string) {
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		handleError(err)
	}
}

// GetFilesAndDirs 获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, dirs, nil
}

// GetAllFiles 获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		fileName := fi.Name()
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fileName)
			GetAllFiles(dirPth + PthSep + fileName)
		} else {
			files = append(files, dirPth+PthSep+fileName)
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}
	return files, nil
}

// filterFile 过滤指定文件名称
func filterFile(fileName string) bool {
	ok := strings.HasSuffix(fileName, ".go")
	if ok {
		return true
	}
	return false
}

// /usr/1.png usr/1.png
// fileNameToObjectName 文件名称转objectName
func fileNameToObjectName(fileName string) string {

	return ""
}
