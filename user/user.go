package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	health int
}

func NewUser() *User {
	return &User{
		health: 100,
	}
} 

func startUILoop(u *User){
	ticker := time.NewTicker(time.Second)
	for {
		fmt.Printf("user healts: %d\n", u.health)
		<-ticker.C
	}
}

func startGameLoop(u *User){
	ticker := time.NewTicker(time.Millisecond * 300)
	for {
		u.health -= rand.Intn(10)
		if u.health <= 0 {
			fmt.Println("Game Over")
			break
		}
		<-ticker.C
	}
}

func main() {
	user := NewUser()
	go startUILoop(user)
	startGameLoop(user)
}
