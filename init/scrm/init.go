package scrm

import (
	"bytes"
	"fmt"
	"gorm.io/gorm"
	databases "moppo.com/gsbt/mysql"
	"strings"
	"time"
)

// 每次查询条数
var SelectLimit int = 5000

func InitWhatsFriend() {
	//
	fmt.Printf("InitWhatsFriend%s", "start")
	// 1kw
	// 13
	//size :=
	partSize := 500000
	start := 0
	for i := 0; i <= 30; i++ {
		end := start + partSize
		fmt.Printf("task i:%d start:%d,end:%d\n", i, start, end)
		go insertTaskPage(0, start, end)
		start = start + partSize
	}
	time.Sleep(200 * time.Minute)

}

func SingleInsert() {
	offset := 0
	for i := 0; i < 40; i++ {
		fmt.Printf("execute task:%d\n", i)
		data := InsertOther(databases.DB, offset)
		if data == nil || len(data) <= 0 {
			continue
		}
		insertPageSize := 500
		// 总的分页数量
		l := len(data)
		page := l / insertPageSize
		startInsertIndex := 0
		endInsertSize := 0
		for i := 0; i <= page; i++ {
			endInsertSize = endInsertSize + insertPageSize
			if endInsertSize > l {
				endInsertSize = l
			}
			insertData := data[startInsertIndex:endInsertSize]
			if insertData == nil || len(insertData) <= 0 {
				break
			}
			BatchSave(databases.DB, insertData)
			startInsertIndex = startInsertIndex + insertPageSize
		}
	}
}

//正常插入
func insertTask(offset, start, end int) {
	for {
		//fmt.Printf("=========task offset:%d start:%d,end:%d\n", offset, start, end)
		us := SelectAllWhatsFriendStatus(offset+1, start, end)
		if us == nil || len(us) <= 0 {
			break
		}
		BatchSave(databases.DB, us)
		offset += SelectLimit
	}
	fmt.Printf("task  over start:%d  end:%d\n", start, end)
}

// insertTaskPage
func insertTaskPage(offset, start, end int) {
	for {
		//fmt.Printf("=========task offset:%d start:%d,end:%d\n", offset, start, end)
		us := SelectAllWhatsFriendStatus(offset+1, start, end)
		if us == nil || len(us) <= 0 {
			break
		}
		offset += SelectLimit
		// 每次insert条数
		insertPageSize := 500
		// 总的分页数量
		l := len(us)
		page := l / insertPageSize
		startInsertIndex := 0
		endInsertSize := 0
		for i := 0; i <= page; i++ {
			endInsertSize = endInsertSize + insertPageSize
			if endInsertSize > l {
				endInsertSize = l
			}
			insertData := us[startInsertIndex:endInsertSize]
			if insertData == nil || len(insertData) <= 0 {
				break
			}
			BatchSave(databases.DB, insertData)
			startInsertIndex = startInsertIndex + insertPageSize
		}
	}
	fmt.Printf("task  over start:%d  end:%d\n", start, end)
}

type TWhatsFriend struct {
	Id                 int    `json:"id" gorm:"column:id;type:bigint;primary_key"`                        //
	TenantId           int    `json:"tenantId" gorm:"column:tenant_id;type:bigint;"`                      //
	WhatsId            string `json:"whatsId" gorm:"column:whats_id;type:character varying;"`             //
	FriendWhatsWhatsId string `json:"whatsId" gorm:"column:friend_whats_id;type:character varying;"`      //
	PushName           string `json:"pushName" gorm:"column:push_name;type:character varying;"`           //
	ChatGroupType      string `json:"chatGroupType" gorm:"column:chatgroup_type;type:int;"`               //
	CreateBy           string `json:"createBy" gorm:"column:create_by;type:int;"`                         //
	CreateAt           string `json:"createAt" gorm:"column:create_at;type:timestamp without time zone;"` //
	DeleteAt           string `json:"deleteAt" gorm:"column:delete_at;type:timestamp without time zone;"` //
}

type TWhatsFriendStatus struct {
	Id            int       `json:"id" gorm:"column:id;type:bigint;primary_key"`                         //
	TenantId      int       `json:"tenantId" gorm:"column:tenant_id;type:bigint;"`                       //
	WhatsFriendId string    `json:"whatsFriendId" gorm:"column:whats_friend_id;type:character varying;"` //
	WhatsId       string    `json:"whatsFriendId" gorm:"column:whats_id;type:character varying;"`        //
	WfId          int       `json:"whatsId" gorm:"column:wf_id;type: bigint;"`                           //
	CreateAt      time.Time `json:"createAt" gorm:"column:create_at;type:timestamp without time zone;"`  //
}

func SelectAllWhatsFriendStatus(offset, start, end int) []TWhatsFriendStatus {
	var wfs []TWhatsFriendStatus
	databases.DB.Offset(offset).Limit(SelectLimit).Where(" id >= ? and id <= ?", start, end).Find(&wfs)
	return wfs
}

func SelectAllWhatsFriendStatusTwo(offset, start, end int) []TWhatsFriendStatus {
	var wfs []TWhatsFriendStatus
	databases.DB.Offset(offset).Not("wf_id", "select id from t_whats_friend").Limit(5).Find(&wfs)
	return wfs
}

func BatchSave(db *gorm.DB, emps []TWhatsFriendStatus) error {
	if emps == nil || len(emps) <= 0 {
		return nil
	}
	var buffer bytes.Buffer
	sql := "insert into `t_whats_friend` (`id`,`tenant_id`,`friend_whats_id`,`whats_id`,`chatgroup_type`,`create_at`,`create_by`) values"
	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}
	for i, e := range emps {
		ct := 0
		// 判断是否群聊
		if strings.Contains(e.WhatsFriendId, "-") {
			ct = 1
		}
		if i == len(emps)-1 {
			buffer.WriteString(fmt.Sprintf("(%d,%d,'%s','%s',%d,'%s',%d);", e.WfId, e.TenantId, e.WhatsFriendId, e.WhatsId, ct, e.CreateAt.Format("2006-01-02 15:04:05.9999"), 101))
		} else {
			buffer.WriteString(fmt.Sprintf("(%d,%d,'%s','%s',%d,'%s',%d),", e.WfId, e.TenantId, e.WhatsFriendId, e.WhatsId, ct, e.CreateAt.Format("2006-01-02 15:04:05.9999"), 101))
		}
	}
	return db.Exec(buffer.String()).Error
}
func InsertOther(db *gorm.DB, offset int) []TWhatsFriendStatus {
	var tfss []TWhatsFriendStatus
	db.Table("t_whats_friend_status").Select("*").Where("wf_id not in( select id from t_whats_friend )").Limit(5000).Offset(offset).Scan(&tfss)
	return tfss
}
