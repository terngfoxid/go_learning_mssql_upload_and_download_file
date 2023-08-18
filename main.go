package main

import (
	"go-mssql-upload/Config"
	_newHandler "go-mssql-upload/delivery"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	//Start DB Connection
	Config.StartConnection()

	r := _newHandler.SetupRouter()
	//running
	r.Run(":8080")
}
