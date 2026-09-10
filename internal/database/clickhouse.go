package database

import (
	"log"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var ClickHouseDB *gorm.DB

func ConnectClickHouse() {
	dsn := "clickhouse://default:@localhost:9000/adali"

	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to clickHouse:", err)
	}

	ClickHouseDB = db
}

