package main

import (
	"github.com/scylladb/gocqlx/qb"
)

func main() {
	qb.Insert("gocqlx_test.person").Columns("first_name", "last_name", "email").ToCql()
}
