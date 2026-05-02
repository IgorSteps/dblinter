package tests

import "database/sql"

type FakeDB struct{}

func (s *FakeDB) SetMaxOpenConns(num int) {}

func testSomething() {
	db, _ := sql.Open("none", "none")
	db.SetMaxOpenConns(10)

	fakeDB := FakeDB{}
	fakeDB.SetMaxOpenConns(99)
}
