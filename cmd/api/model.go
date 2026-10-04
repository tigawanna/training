package main

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID   int64  `bun:",pk,autoincrement" json:"id"`
	Name string `bun:",notnull" json:"name"`
}
