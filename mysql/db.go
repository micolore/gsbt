package databases

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

var LocalDbName = "test"
var LocalDbPWD = "root"
var LocalDbUserName = "root"
var LocalDbIP = "localhost"
var LocalDbPort = 3306

var TestFiveDbName = "wscrm"
var TestFiveDbPWD = "123Moppo!@#"
var TestFiveDbUserName = "root"
var TestFiveDbIP = "192.168.1.5"
var TestFiveDbPort = 3306

var TestFourDbName = "scrm_bus"
var TestFourDbPWD = "mysql2Moppo123#"
var TestFourDbUserName = "root"
var TestFourDbIP = "192.168.1.4"
var TestFourDbPort = 34521

func getProfile(profileType int) string {
	dbURL := ""
	switch profileType {
	// localhost
	case 1:
		dbURL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", LocalDbUserName, LocalDbPWD, LocalDbIP, LocalDbPort, LocalDbName)
	// 4
	case 2:
		dbURL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", TestFourDbUserName, TestFourDbPWD, TestFourDbIP, TestFourDbPort, TestFourDbName)
	// 5
	case 3:
		dbURL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", TestFiveDbUserName, TestFiveDbPWD, TestFiveDbIP, TestFiveDbPort, TestFiveDbName)
	}
	return dbURL
}

func init() {
	dbURL := getProfile(1)
	var err error
	DB, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	if DB != nil {
		fmt.Printf("db conn success!\n")
	}
}
