package Config

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

var server = "127.0.0.1"
var port = 1433
var user = "admin"
var password = "12345678"
var database = "test_go_ms_sql"

func GetConStr() string {
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
}

func StartConnection() {
	// Build connection string
	connString := GetConStr()

	var err error

	// Create connection
	DB, err = gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		panic("database connection failed : " + err.Error())
	}
	fmt.Printf("Connected!\n")
}
