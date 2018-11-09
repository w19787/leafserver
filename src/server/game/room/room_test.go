package room

import (
	"server/game/player"
	"testing"
)

func TestNewGameRoom(t *testing.T) {
	g := NewGameRoom(2)
	if g == nil {
		t.Error("create failed")
	}

	if g.Size != 2 {
		t.Error("size is incorrect")
	}

	if g.Status != STATUS_OPEN {
		t.Error("status is not OPEN")
	}

	if len(g.Players) != 0 {
		t.Error("players len is incorrect")
	}

	if cap(g.Players) != 2 {
		t.Error("players cap is incorrect")
	}

	go func() {
		m := Msg{player.Player{Name: "alex", Level: 2}, MSG_JOIN}
		g.MsgChan <- m
	}()

	rm := <-g.MsgChan
	if rm.T != MSG_JOIN {
		t.Error("receive msg type error")
	}

	if rm.P.Name != "alex" {
		t.Error("player name error")
	}

	if rm.P.Level != 2 {
		t.Error("player level error")
	}
}

func TestOnJoin(t *testing.T) {
	g := NewGameRoom(2)
	p := player.Player{Name: "alex", Level: 2}
	result := g.onJoin(p)
	if !result {
		t.Error("onJoin error")
	}

	if len(g.Players) != 1 || cap(g.Players) != 2 {
		t.Error("players cap error", cap(g.Players))
	}

	result = g.onJoin(p)
	if !result {
		t.Error("onJoin error")
	}

	if len(g.Players) != 2 || cap(g.Players) != 2 {
		t.Error("players cap error", cap(g.Players))
	}

	result = g.onJoin(p)
	if result {
		t.Error("onJoin error")
	}

	if len(g.Players) != 2 || cap(g.Players) != 2 {
		t.Error("players cap error", cap(g.Players))
	}
}

func TestOnLeave(t *testing.T) {
	g := NewGameRoom(2)
	p := player.Player{Name: "alex", Level: 2}
	result := g.onJoin(p)
	if !result {
		t.Error("onJoin error")
	}

	g.onLeave(p)
	if len(g.Players) != 0 || cap(g.Players) != 2 {
		t.Error("players cap error", cap(g.Players))
	}

	g.onLeave(p)
	if len(g.Players) != 0 || cap(g.Players) != 2 {
		t.Error("players cap error", cap(g.Players))
	}
}

func TestProcessOnJoin(t *testing.T) {
	g := NewGameRoom(2)

	go g.Process()
	m := Msg{player.Player{Name: "alex", Level: 2}, MSG_JOIN}
	g.MsgChan <- m

	if len(g.Players) != 1 || cap(g.Players) != 2 {
		t.Error("players cap error", cap(g.Players))
	}

	if g.Players[0].Name != "alex" {
		t.Error("Name is error")
	}
}

func TestProcessOnLeave(t *testing.T) {
	g := NewGameRoom(2)
	p := player.Player{Name: "alex", Level: 2}
	g.onJoin(p)

	if len(g.Players) != 1 || cap(g.Players) != 2 {
		t.Error("players cap error", cap(g.Players))
	}

	go g.Process()
	m := Msg{player.Player{Name: "alex", Level: 2}, MSG_LEAVE}
	g.MsgChan <- m

	if len(g.Players) != 0 || cap(g.Players) != 2 {
		t.Error("players cap error", cap(g.Players))
	}
}
