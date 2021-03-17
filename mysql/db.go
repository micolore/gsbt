package databases

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

// LocalHost
func init() {
	dbUserName := "root"
	dbPWD := "root"
	dbIP := "localhost"
	dbPort := 3306
	dbName := "test" //practice_questions
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbUserName, dbPWD, dbIP, dbPort, dbName)
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
		fmt.Printf("db != nil\n")
	}
}

func initFive() {
	dbUserName := "root"
	dbPWD := "123Moppo!@#"
	dbIP := "192.168.1.5"
	dbPort := 3306
	dbName := "zhu"
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbUserName, dbPWD, dbIP, dbPort, dbName)
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
		fmt.Printf("db != nil\n")
	}
}
func initFiveOff() {
	dbUserName := "root"
	dbPWD := "123Moppo!@#"
	dbIP := "192.168.1.5"
	dbPort := 3306
	dbName := "wscrm"
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbUserName, dbPWD, dbIP, dbPort, dbName)
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
		fmt.Printf("db != nil\n")
	}
}
