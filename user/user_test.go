package main

import "testing"

func TestGame(t *testing.T){
	user := NewUser()
	go startUILoop(user)
	startGameLoop(user)
}