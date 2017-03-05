package robot

import (
	"github.com/kidoman/embd"
	"github.com/kidoman/embd/controller/pca9685"
	_ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
)

type ServoController struct {
	pca  *pca9685.PCA9685
}

func NewServoController(addr byte) *ServoController {
	bus := embd.NewI2CBus(1)
	return &ServoController{
		pca: pca9685.New(bus, addr),
	}
}

func (c *ServoController) Servo(channel int) *Servo {
	return &Servo{
		ctrl: c,
		channel: channel,
	}
}

func (c *ServoController) Close() error {
	return c.pca.Close()
}

func (c *ServoController) Sleep() error {
	return c.pca.Sleep()
}

func (c *ServoController) Wake() error {
	return c.pca.Wake()
}
