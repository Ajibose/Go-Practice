package enums

type PlayerState int
type PlayerAction int

const (
	Stopped PlayerState = iota
	Playing
	Paused
)

const (
	PressPlay PlayerAction = iota
	PressPause
	PressStop
)

/* func (ps PlayerState) String() string {
	var PlayerStateName = map[PlayerState]string{
		Stopped: "Music stopped",
		Paused:  "Music paused",
		Playing: "Music started",
	}

	return PlayerStateName[ps]
} */

func OperatePlayer(current PlayerState, action PlayerAction) (PlayerState, string) {
	switch {
	case current == Stopped && action == PressPlay:
		return Playing, "Music Started"
	case current == Playing && action == PressPause:
		return Paused, "Music Paused"
	case current == Playing && action == PressStop:
		return Stopped, "Music stopped"
	case current == Paused && action == PressPlay:
		return Playing, "Music resumed"
	case current == Paused && action == PressStop:
		return Stopped, "Music stopped"
	default:
		return current, "Invalid action for current state"
	}
}
