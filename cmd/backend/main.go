package main

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("后端")
}
