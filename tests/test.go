package tests

import "database/sql"

func testSomething() {
	db, _ := sql.Open("none", "none")
	db.SetMaxOpenConns(10)
}
