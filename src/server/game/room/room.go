func (GR *GameRoom) Process() {
	// 处理游戏过程结束/正常/异常
	defer func() {
		GR.End()
		if e := recover(); e != nil {
			panic(e)
		}
	}()
	GR.Playing = true
	gameStartTime := conf.GameConfig.StartTime
	robotTime := conf.GameConfig.RobotTime
Playing:
	for {
		select {
		case m := <-GR.MsgChan:
			player := m.P
			switch m.Mtype {
			case MSG_JOIN:
				GR.onJoin(player)
			case MSG_ANSWER:
				GR.onAnswer(player, m.Answer, m.Pm)
			}
		case wf := <-GR.WaitChan:
			switch wf {
			case WAIT_GAME:
				// 游戏开始倒计时，补充一部分AI
				joinRobots(GR, utils.RandIntFrom2(18, 25))
				// 等待是否补充足房间
				time.AfterFunc(time.Second*time.Duration(robotTime), func() {
					GR.WaitChan <- ON_RO_JOIN
				})
			case ON_RO_JOIN:
				// 不足AI
				joinRobots(GR, 0)
				// 剩余倒计时
				time.AfterFunc(time.Second*time.Duration(gameStartTime-robotTime), func() {
					GR.WaitChan <- ON_GAME_START
				})
			case ON_GAME_START:
				// 游戏开始
				GR.Start()
			case ON_RO_ANS_1:
				// 机器人第一次回答
				helpRobotAnswer(GR, true, false)
				time.AfterFunc(time.Second*time.Duration(1), func() {
					GR.WaitChan <- ON_RO_ANS_2
				})
			case ON_RO_ANS_2:
				// 机器人第二次回答
				helpRobotAnswer(GR, false, false)
				time.AfterFunc(time.Second*time.Duration(1), func() {
					GR.WaitChan <- ON_RO_ANS_3
				})
			case ON_RO_ANS_3:
				helpRobotAnswer(GR, false, true)
				// 倒计时回答结束
				time.AfterFunc(time.Second*time.Duration(GR.AnswerTimeout-4), func() {
					GR.WaitChan <- ON_ANSWER_END
				})
			case ON_ANSWER_END: // 回答结束，结算
				GR.onAnswerOver()
			case ON_SEND_QUESTION: // 出题
				GR.sendQuestion()
			case MSG_END_GAME: // 游戏结束
				//break game
				break Playing
				// goto END
			}
		}
	}
	//END:
	log.Debug("Game goroutine over ->%s", GR.Uuid)
}