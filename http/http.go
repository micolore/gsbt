package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	for i := 0; i <= 10000; i++ {
		  trunslate()
	}
	time.Sleep(time.Hour * 2)
}

func tls() {
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//client := &http.Client{Transport: tr}

	resp, err := http.Post("https://api-scrm.rupiahcepatweb.com/wscrm-bus-api/api/authToken/whatsappContact", "application/json", nil)
	//resp, err := http.Post("https://api-scrm.rupiahcepatweb.com/api/proxy/home/messageCount", "application/json", nil)
	//resp, err := http.Post("http://baidu.com", "application/json", nil)

	if err != nil {
		panic(err)
		fmt.Printf("post request failed, err:[%s]", err.Error())
		return
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp status code:[%d]\n", resp.StatusCode)
	fmt.Printf("resp body data:[%s]\n", string(bodyContent))

}

func trunslate() {
	resp, err := http.Get("http://192.168.1.6:10001/test/test?tenantId=405&translateByteLength=10")
	if err != nil {
		panic(err)
		fmt.Printf("post request failed, err:[%s]", err.Error())
		return
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("resp status code:[%d]\n", resp.StatusCode)
	fmt.Printf("resp body data:[%s]\n", string(bodyContent))
}
