package tests

import (
	"database/sql"
	"fmt"
)

func testSomething() {
	db, _ := sql.Open("none", "none")
	fmt.Printf("db: %v\n", db)
}
