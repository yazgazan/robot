package robot

type Direction byte

const (
	Stop     Direction = 0x00
	Forward  Direction = 0x01
	Backward Direction = 0x02
)

func (d Direction) String() string {
	switch d {
	case 0x00:
		return "stop"
	case 0x01:
		return "forward"
	case 0x02:
		return "backward"
	default:
		return "<unknown>"
	}
}
