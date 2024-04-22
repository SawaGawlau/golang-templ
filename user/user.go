package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type User struct {
	mu sync.RWMutex // read-write mutex
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
		u.mu.RLock()
		fmt.Printf("user healts: %d\n", u.health)
		u.mu.RUnlock()
		<-ticker.C
	}
}

func startGameLoop(u *User){
	ticker := time.NewTicker(time.Millisecond * 300)
	for {
		u.mu.Lock() // lock read & write
		u.health -= rand.Intn(10)
		if u.health <= 0 {
			fmt.Println("Game Over")
			break
		}
		u.mu.Unlock() // unlock read & write
		<-ticker.C
	}
}

func main() {
	user := NewUser()
	go startUILoop(user)
	startGameLoop(user)
}
