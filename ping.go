package robot

type Ping struct {
	i2cHandler *I2C
	addr       byte
}

func NewPing(handler *I2C, addr byte) *Ping {
	return &Ping{
		i2cHandler: handler,
		addr:       addr,
	}
}

func (p *Ping) Read() (int, error) {
	b, err := p.i2cHandler.ReadByte(p.addr)

	return int(b), err
}
