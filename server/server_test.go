package server

import (
	"fmt"
	"testing"
)
type Player struct {
	Name string
}

type GameState struct {
	// lock sync.RWMutex
	players []*Player 
	msgch chan any // message channel
}

func (g *GameState) receive(msg any){
	g.msgch <- msg

}
func NewGameState() *GameState {
	g := &GameState{
		players: []*Player{},
		msgch: make(chan any, 10),
	}
	go g.loop()
	return g
}

type Server struct {
	gameState *GameState
}

func (g *GameState) loop() {
	for msg := range g.msgch {
		// handle message
		g.handleMessage(msg)
	}

}

func (g *GameState) handleMessage(message any){
	switch msg := message.(type) {
		case *Player:
			g.addPlayer(msg)
		default: 
			panic("invalid message type")
		}
}

func (g *GameState) addPlayer(p *Player) {
	// g.lock.Lock()
	g.players = append(g.players, p)
	// g.lock.Unlock()

	fmt.Println("adding platyer", p.Name)

}
func NewServer() *Server {
	return &Server{
		gameState: NewGameState(),
	}
}

func (s *Server) handleNewPlayer(player *Player) {
	s.gameState.receive(player)
}

func TestServer(t *testing.T) {
	server := NewServer()

	for i := 0; i < 10; i++ {
		player := &Player{
			Name: fmt.Sprintf("Player_%d", i),
		}
	go server.handleNewPlayer(player)
	}
}