package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
)

// elasticsearch
func main() {
	gets()
}

type MessageRecord struct {
	QuotedMsgSvrid string `json:"quoted_msg_svrid"`
	TenantId       string `json:"tenant_id"`
	DeleteTime     int    `json:"delete_time"`
	MsgType        string `json:"msgType"`
	CreateTime     string `json:"create_time"`
	ActionType     string `json:"action_type"`
	SenderWhatsId  string `json:"senderWhatsId"`
	IsChatGroup    string `json:"isChatGroup"`
	WhatsTime      string `json:"whatsTime"`
	OriginType     string `json:"origin_type"`
	Content        string `json:"content"`
	ChatType       string `json:"chat_type"`
	AccountId      string `json:"accountId"`
	SendTime       string `json:"send_time"`
	ContentType    string `json:"content_type"`
	UserId         string `json:"user_id"`
	CurrentWhatsId string `json:"current_whats_id"`
	IsSend         string `json:"isSend"`
	Id             string `json:"id"`
	FriendWhatsId  string `json:"friend_whats_id"`
	MsgId          string `json:"msg_id"`
}
var client *elastic.Client
var host = "http://192.168.1.4:9100"
//初始化
func init() {
	//errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error
	//这个地方有个小坑 不加上elastic.SetSniff(false) 会连接不上
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	_, _, err = client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	_, err = client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Elasticsearch version %s\n", esversion)

}

/*下面是简单的CURD*/

//查找
func gets() {
	//通过id查找
	get1, err := client.Get().Index("wscrm_message_new").Id("2").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
		var bb MessageRecord
		err := json.Unmarshal(get1.Source, &bb)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(bb.Content)
		fmt.Println(string(get1.Source))
	}

}

//
//删除
func delete() {

	res, err := client.Delete().Index("megacorp").
		Type("employee").
		Id("1").
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

//
////搜索
func query() {
	var res *elastic.SearchResult
	var err error
	//取所有
	res, err = client.Search("megacorp").Type("employee").Do(context.Background())
	printEmployee(res, err)

	//字段相等
	q := elastic.NewQueryStringQuery("last_name:Smith")
	res, err = client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	printEmployee(res, err)

	//条件查询
	//年龄大于30岁的
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	res, err = client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
	printEmployee(res, err)

	//短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
	res, err = client.Search("megacorp").Type("employee").Query(matchPhraseQuery).Do(context.Background())
	printEmployee(res, err)

	//分析 interests
	aggs := elastic.NewTermsAggregation().Field("interests")
	res, err = client.Search("megacorp").Type("employee").Aggregation("all_interests", aggs).Do(context.Background())
	printEmployee(res, err)

}

//
////简单分页
func list(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err := client.Search("megacorp").
		Type("employee").
		Size(size).
		From((page - 1) * size).
		Do(context.Background())
	printEmployee(res, err)

}

//
//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ MessageRecord
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(MessageRecord)
		fmt.Printf("%#v\n", t)
	}
}
