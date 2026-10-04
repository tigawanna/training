package main

import "sync"

type User struct {
	Name string `json:"name"`
}

var (
	mutex sync.RWMutex
	nextID = 2
	Users  = map[int]User{
		1: {Name: "User 1"},
	}
)
