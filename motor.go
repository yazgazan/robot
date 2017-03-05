package robot

type Motor struct {
	Name       string
	i2cHandler *I2C
	dirAddr    byte
	speedAddr  byte
}

func NewMotor(name string, handler *I2C, dirAddr byte, speedAddr byte) *Motor {
	return &Motor{
		Name:       name,
		i2cHandler: handler,
		dirAddr:    dirAddr,
		speedAddr:  speedAddr,
	}
}

func (m *Motor) GetDir() (Direction, error) {
	b, err := m.i2cHandler.ReadByte(m.dirAddr)

	return Direction(b), err
}

func (m *Motor) GetSpeed() (int, error) {
	b, err := m.i2cHandler.ReadByte(m.speedAddr)

	return int(b), err
}

func (m *Motor) SetDir(dir Direction) error {
	return m.i2cHandler.WriteByte(m.dirAddr, byte(dir))
}

func (m *Motor) SetSpeed(speed int) error {
	return m.i2cHandler.WriteByte(m.speedAddr, byte(speed))
}
