package tests

import "database/sql"

type FakeDB struct{}

func (s *FakeDB) SetMaxOpenConns(num int) {}

func (s *FakeDB) DoSomething(boo int) {}

func testSomething() {
	db, _ := sql.Open("none", "none")
	db.SetMaxOpenConns(11)
	db.SetMaxOpenConns(12)
	db.SetMaxOpenConns(13)

	fakeDB := FakeDB{}
	fakeDB.SetMaxOpenConns(99)
	fakeDB.DoSomething(123)
}
