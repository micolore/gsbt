package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	for i := 0; i < 500; i++ {
		execTask()
		time.Sleep(5 * time.Second)
	}

}

func execTask() {
	resp, err := http.Get("http://wascrm.socialepoch.com/wscrm-bus-api/getFriendInfo?friendWhatsId=8618827017303-1635872577@g.us&whatsId=8618217331413@c.us&pushName=qwerdc&tenantId=82&accountId=1114&departmentId=81")

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
