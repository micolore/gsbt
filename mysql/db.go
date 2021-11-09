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

func getProfile(profileType int) string {
	dbURL := ""
	switch profileType {
	// localhost
	case 1:
		dbURL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", LocalDbUserName, LocalDbPWD, LocalDbIP, LocalDbPort, LocalDbName)
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
