package robot

import (
	"errors"
	"golang.org/x/exp/io/i2c"
)

type I2C struct {
	handler *i2c.Device
}

func OpenI2C(device string, addr int) (*I2C, error) {
	handler, err := i2c.Open(&i2c.Devfs{Dev: device}, addr)

	if err != nil {
		return nil, err
	}

	return &I2C{
		handler: handler,
	}, nil
}

func (h *I2C) ReadByte(addr byte) (byte, error) {
	b := make([]byte, 1)
	err := h.handler.ReadReg(addr, b)

	if err != nil {
		return 0, err
	}

	return b[0], nil
}

func (h *I2C) WriteByte(addr byte, b byte) error {
	return h.handler.WriteReg(addr, []byte{b})
}

func (h *I2C) Close() {
	if h.handler == nil {
		panic(errors.New("Cannot close nil handler"))
	}
	h.handler.Close()
}
