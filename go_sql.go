package main

import (
	"fmt"
)
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func main(){
	db, err := sql.Open("mysql","jetty:123456@tcp(127.0.0.1:3306)/tiku")
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Connection Established")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}

