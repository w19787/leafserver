package room

import (
	"server/game/player"
)

type MsgType int

const (
	MSG_JOIN MsgType = iota
	MSG_WAIT_GAME
	MSG_ON_GAME_START
	MSG_ON_RO_JOIN
	MSG_END_GAME
	MSG_LEAVE
)

type RoomStatus int

const (
	STATUS_OPEN RoomStatus = iota
	STATUS_PLAYING
	STATUS_CLOSE
)

type Msg struct {
	P player.Player
	T MsgType
}

type GameRoom struct {
	Size    int
	MsgChan chan Msg
	Status  RoomStatus
	Uuid    int
	Players []player.Player
}

func NewGameRoom(size int) *GameRoom {
	gr := new(GameRoom)
	gr.Size = size
	gr.MsgChan = make(chan Msg)
	gr.Status = STATUS_OPEN
	gr.Players = make([]player.Player, 0, size)

	return gr
}

func (gr *GameRoom) onJoin(p player.Player) bool {
	if len(gr.Players) < gr.Size {
		gr.Players = append(gr.Players, p)
		return true
	}

	return false
}

func (gr *GameRoom) onLeave(p player.Player) {
	for i := 0; i < len(gr.Players); i++ {
		if gr.Players[i] == p {
			gr.Players = append(gr.Players[:i], gr.Players[i+1:]...)
			break
		}
	}
}

func (gr *GameRoom) startGame() {
	gr.Status = STATUS_PLAYING
}

func (gr *GameRoom) endGame() {
	gr.Status = STATUS_CLOSE

}

func (gr *GameRoom) Process() {
	// 处理游戏过程结束/正常/异常
	defer func() {
		gr.endGame()
		if e := recover(); e != nil {
			panic(e)
		}
	}()

	for {
		select {
		case m := <-gr.MsgChan:
			player := m.P
			switch m.T {
			case MSG_JOIN:
				gr.onJoin(player)
			case MSG_WAIT_GAME:
				// 游戏开始倒计时，补充一部分AI
				// joinRobots(gr, utils.RandIntFrom2(18, 25))
				// // 等待是否补充足房间
				// time.AfterFunc(time.Second*time.Duration(robotTime), func() {
				// 	gr.MsgChan <- ON_RO_JOIN
				// })
			case MSG_ON_RO_JOIN:
				// 不足AI
				// joinRobots(gr, 0)
				// // 剩余倒计时
				// time.AfterFunc(time.Second*time.Duration(gameStartTime-robotTime), func() {
				// 	gr.MsgChan <- ON_GAME_START
				// })
			case MSG_ON_GAME_START:
				// 游戏开始
				gr.startGame()

			case MSG_LEAVE:
				gr.onLeave(player)

			case MSG_END_GAME: // 游戏结束
				//break game
				break
			}
		}
	}

}
