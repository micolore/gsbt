package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type TPv struct {
	Id                int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	TenantId          int       `gorm:"column:tenant_id" db:"tenant_id" json:"tenant_id" form:"tenant_id"`                                         //公司id
	UserId            int       `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`                                                 //用户id
	UserName          string    `gorm:"column:user_name" db:"user_name" json:"user_name" form:"user_name"`                                         //用户账号  有可能是游客
	PhoneNum          string    `gorm:"column:phone_num" db:"phone_num" json:"phone_num" form:"phone_num"`                                         //用户手机号
	CookieUserId      string    `gorm:"column:cookie_user_id" db:"cookie_user_id" json:"cookie_user_id" form:"cookie_user_id"`                     //cookie 用户 id
	UserActionType    string    `gorm:"column:user_action_type" db:"user_action_type" json:"user_action_type" form:"user_action_type"`             //用户行为类型（多个系统不重复的行为类型）
	PageRemainTime    int       `gorm:"column:page_remain_time" db:"page_remain_time" json:"page_remain_time" form:"page_remain_time"`             //页面停留时长
	PageLocation      int       `gorm:"column:page_location" db:"page_location" json:"page_location" form:"page_location"`                         //页面停留位置 （1-5）
	DataSource        string    `gorm:"column:data_source" db:"data_source" json:"data_source" form:"data_source"`                                 //数据来源  1 安卓客户端 2、pc web端 3、mobile web端 4、pc客户端
	ClientVersionName string    `gorm:"column:client_version_name" db:"client_version_name" json:"client_version_name" form:"client_version_name"` //1.12.1
	ClientVersionCode string    `gorm:"column:client_version_code" db:"client_version_code" json:"client_version_code" form:"client_version_code"` //111200110
	Uri               string    `gorm:"column:uri" db:"uri" json:"uri" form:"uri"`                                                                 //请求URI
	Channel           string    `gorm:"column:channel" db:"channel" json:"channel" form:"channel"`                                                 //渠道号
	MPhoneModel       string    `gorm:"column:m_phone_model" db:"m_phone_model" json:"m_phone_model" form:"m_phone_model"`                         //机型
	MAndroidVersion   string    `gorm:"column:m_android_version" db:"m_android_version" json:"m_android_version" form:"m_android_version"`         //android版本
	MResolution       string    `gorm:"column:m_resolution" db:"m_resolution" json:"m_resolution" form:"m_resolution"`                             //分辨率
	MSerial           string    `gorm:"column:m_serial" db:"m_serial" json:"m_serial" form:"m_serial"`                                             //android序列号
	MAppLocale        string    `gorm:"column:m_app_locale" db:"m_app_locale" json:"m_app_locale" form:"m_app_locale"`                             //app语言选项
	MVersionCode      string    `gorm:"column:m_version_code" db:"m_version_code" json:"m_version_code" form:"m_version_code"`                     //构件号
	MOaid             string    `gorm:"column:m_oaid" db:"m_oaid" json:"m_oaid" form:"m_oaid"`                                                     //oaid
	MImei             string    `gorm:"column:m_imei" db:"m_imei" json:"m_imei" form:"m_imei"`                                                     //imei
	MAndroidId        string    `gorm:"column:m_android_id" db:"m_android_id" json:"m_android_id" form:"m_android_id"`                             //android id
	MImsi             string    `gorm:"column:m_imsi" db:"m_imsi" json:"m_imsi" form:"m_imsi"`                                                     //imsi
	MIdfa             string    `gorm:"column:m_idfa" db:"m_idfa" json:"m_idfa" form:"m_idfa"`                                                     //idfa iphone
	MIdfv             string    `gorm:"column:m_idfv" db:"m_idfv" json:"m_idfv" form:"m_idfv"`                                                     //idfy iphone
	MUdid             string    `gorm:"column:m_udid" db:"m_udid" json:"m_udid" form:"m_udid"`                                                     //udid iphone
	MOpenid           string    `gorm:"column:m_openid" db:"m_openid" json:"m_openid" form:"m_openid"`                                             //openid iphone
	Country           string    `gorm:"column:country" db:"country" json:"country" form:"country"`                                                 //国家
	Province          string    `gorm:"column:province" db:"province" json:"province" form:"province"`                                             //省
	Ip                string    `gorm:"column:ip" db:"ip" json:"ip" form:"ip"`                                                                     //ip
	ProxyIp           string    `gorm:"column:proxy_ip" db:"proxy_ip" json:"proxy_ip" form:"proxy_ip"`                                             //proxy ip
	Network           string    `gorm:"column:network" db:"network" json:"network" form:"network"`                                                 //网络
	Locale            string    `gorm:"column:locale" db:"locale" json:"locale" form:"locale"`                                                     //语言
	Longitude         string    `gorm:"column:longitude" db:"longitude" json:"longitude" form:"longitude"`                                         //经度
	Latitude          string    `gorm:"column:latitude" db:"latitude" json:"latitude" form:"latitude"`                                             //维度
	DeviceId          string    `gorm:"column:device_id" db:"device_id" json:"device_id" form:"device_id"`                                         //Google advertising ID advertising ID===  device_id  com.google.android.gms.ads.identifier
	Vendor            string    `gorm:"column:vendor" db:"vendor" json:"vendor" form:"vendor"`                                                     //制造商
	Mac               string    `gorm:"column:mac" db:"mac" json:"mac" form:"mac"`                                                                 //mac地址
	UserAgent         string    `gorm:"column:user_agent" db:"user_agent" json:"user_agent" form:"user_agent"`                                     //user_agent
	Referrer          string    `gorm:"column:referrer" db:"referrer" json:"referrer" form:"referrer"`                                             //上一页面
	GpReferrer        string    `gorm:"column:gp_referrer" db:"gp_referrer" json:"gp_referrer" form:"gp_referrer"`                                 //gp referrer
	WebUrl            string    `gorm:"column:web_url" db:"web_url" json:"web_url" form:"web_url"`                                                 //当前页面
	Cookie            string    `gorm:"column:cookie" db:"cookie" json:"cookie" form:"cookie"`                                                     //cookie信息
	CreateTime        time.Time `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`                                 //日期
	RequestParams     string    `gorm:"column:request_params" db:"request_params" json:"request_params" form:"request_params"`                     //请求参数
}

func PvInsertSql() string {
	tPv := TPv{}
	tPvSql := GetSructField(tPv)
	sql := fmt.Sprintf("INSERT INTO t_pv (%s)", tPvSql)
	return sql
}

func GetSructField(s interface{}) string {

	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	fieldNum := t.NumField()
	fieldSql := ""
	for i := 0; i < fieldNum; i++ {
		fieldName := t.Field(i).Name
		if i == fieldNum-1 {
			fieldSql += fieldName
			break
		} else {
			fieldSql += fieldName + ","
		}
	}
	return fieldSql
}
func example() error {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"8.214.3.31:19000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		//Debug:           true,
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
	})
	if err != nil {
		return err
	}
	ctx := clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
		"max_block_size": 10,
	}), clickhouse.WithProgress(func(p *clickhouse.Progress) {
		fmt.Println("progress: ", p)
	}))
	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Catch exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return err
	}
	if err := conn.Exec(ctx, `DROP TABLE IF EXISTS example`); err != nil {
		return err
	}
	err = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS example (
			Col1 UInt8,
			Col2 String,
			Col3 DateTime
		) engine=Memory
	`)
	if err != nil {
		return err
	}
	batch, err := conn.PrepareBatch(ctx, "INSERT INTO example (Col1, Col2, Col3)")
	if err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		if err := batch.Append(uint8(i), fmt.Sprintf("value_%d", i), time.Now()); err != nil {
			return err
		}
	}
	if err := batch.Send(); err != nil {
		return err
	}

	rows, err := conn.Query(ctx, "SELECT Col1, Col2, Col3 FROM example WHERE Col1 >= $1 AND Col2 <> $2 AND Col3 <= $3", 0, "xxx", time.Now())
	if err != nil {
		return err
	}
	for rows.Next() {
		var (
			col1 uint8
			col2 string
			col3 time.Time
		)
		if err := rows.Scan(&col1, &col2, &col3); err != nil {
			return err
		}
		fmt.Printf("row: col1=%d, col2=%s, col3=%s\n", col1, col2, col3)
	}
	rows.Close()
	return rows.Err()
}

// 初始化亿万的数据
func initMillionData() {

	sql := PvInsertSql()
	fmt.Println(sql)
}

func main() {
	if err := example(); err != nil {
		log.Fatal(err)
	}
	initMillionData()
}
