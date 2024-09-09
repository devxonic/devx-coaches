package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"

	execapi "github.com/ibad69/go-devxcoaches/src/api"
)

func main() {
	fmt.Print("devxcoahes live on golang")
	uri := "postgresql://postgres:yourpassword@137.59.222.200:5432/devxcoaches"
	execapi.Start(&execapi.ApiConfig{Dburi: uri, Port: "5100"})
}
