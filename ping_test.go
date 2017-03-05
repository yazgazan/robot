package robot

import (
	"testing"
)

func TestPing(t *testing.T) {
	var err error

	handler, err := OpenI2C("/dev/i2c-1", 0x08)
	if err != nil {
		t.Logf("error opening i2c: %s", err)
		t.Fail()
		return
	}
	defer handler.Close()

	ping := NewPing(handler, 10)

	dist, err := ping.Read()
	if err != nil {
		t.Logf("error reading distance: %s", err)
		t.Fail()
		return
	}

	t.Logf("read distance %d", dist)
}
