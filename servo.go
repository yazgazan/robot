package robot

import (
	"github.com/kidoman/embd/util"
)

const (
	pwmControlPoints = 4096
)

type Servo struct{
	ctrl *ServoController
	channel int
}

func (s *Servo) Write(value byte) error {
	offTime := util.Map(int64(value), 0, 255, 0, pwmControlPoints-1)
	return s.ctrl.pca.SetPwm(s.channel, 0, int(offTime))
}

func (s *Servo) WriteAngle(angle int) error {
	value := 50 + angle * 200 / 180

	return s.Write(byte(value))
}
