package mission

type MissionState int
const(
	MISSION_STATE_OPEN MissionState = itoa
	MISSION_STATE_ACTIVE
	MISSION_STATU_COMPLETED
)


struct Mission{
	Name string
	State MissionState
}