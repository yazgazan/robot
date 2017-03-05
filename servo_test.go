package robot

import (
	"time"
	"testing"
)

func TestServo(t *testing.T) {
	ctrl := NewServoController(0x40)
	servo := ctrl.Servo(0)
	defer func() {
		if err := ctrl.Close(); err != nil {
			t.Logf("error closing ServoController: %s", err)
			t.Fail()
		}
	}()

	testServo := func (angle int) {
		t.Logf("writing angle %d to servo", angle)
		if err := servo.WriteAngle(angle); err != nil {
			t.Logf("error writing to servo: %s", err)
			t.Fail()
		}
	}

	// Servo.Write
	// left-most: 50
	// right-most: 250
	// center: 150
	// Servo.WriteAngle
	// left-most: 0
	// right-most: 180
	// center: 90
	for i := 0; i < 180; i += 15 {
		testServo(i)
		<-time.After(1 * time.Second)
	}
	testServo(90)
	<-time.After(2 * time.Second)
	if err := ctrl.Sleep(); err != nil {
		t.Logf("error puting controller to sleep: %s", err)
		t.Fail()
	}
}
