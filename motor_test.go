package robot

import (
	"testing"
	"time"
)

func TestMotor(t *testing.T) {
	var err error

	handler, err := OpenI2C("/dev/i2c-1", 0x08)
	if err != nil {
		t.Logf("error opening i2c: %s", err)
		t.Fail()
		return
	}
	defer handler.Close()

	motorA := NewMotor("A", handler, 31, 30)
	motorB := NewMotor("B", handler, 33, 32)

	motorStatus := func(motor *Motor) {
		dir, err := motor.GetDir()
		if err != nil {
			t.Logf("error reading motor %s direction: %s", motor.Name, err)
			t.Fail()
			return
		}
		speed, err := motor.GetSpeed()
		if err != nil {
			t.Logf("error reading motor %s speed: %s", motor.Name, err)
			t.Fail()
			return
		}

		t.Logf("motor %s: dir=%s; speed=%d", motor.Name, dir, speed)
	}

	motorStance := func(motor *Motor, dir Direction, speed int) {
		if dir != Stop {
			err = motor.SetDir(dir)
			if err != nil {
				t.Logf(
					"error setting motor %s direction to %s: %s",
					motor.Name,
					dir,
					err,
				)
				t.Fail()
				return
			}
		}

		err = motor.SetSpeed(speed)
		if err != nil {
			t.Logf(
				"error setting motor %s speed to %d: %s",
				motor.Name,
				speed,
				err,
			)
			t.Fail()
			return
		}
	}

	motorStatus(motorA)
	motorStatus(motorB)
	motorStance(motorA, Forward, 100)
	motorStance(motorB, Forward, 100)

	<-time.After(1 * time.Second)
	motorStatus(motorA)
	motorStatus(motorB)
	motorStance(motorA, Stop, 0)
	motorStance(motorB, Stop, 0)

	<-time.After(1 * time.Second)
	motorStatus(motorA)
	motorStatus(motorB)
	motorStance(motorA, Backward, 100)
	motorStance(motorB, Backward, 100)

	<-time.After(1 * time.Second)
	motorStatus(motorA)
	motorStatus(motorB)
	motorStance(motorA, Stop, 0)
	motorStance(motorB, Stop, 0)
}
