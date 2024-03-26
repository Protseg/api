package database

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

var (
	dbUser     string = os.Getenv("DB_USER")
	dbPassword string = os.Getenv("DB_PASSWORD")
	dbName     string = os.Getenv("DB_NAME")
	dbPort     string = os.Getenv("DB_PORT")
	dbIp       string = os.Getenv("DB_IP")
	DBConn     *gorm.DB
	ConnString string = fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbIp,
		dbPort,
		dbName,
	)
)
