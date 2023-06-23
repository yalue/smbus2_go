package smbus2_go

import (
	"testing"
)

func TestOpenBus(t *testing.T) {
	busID := 0
	bus, e := NewSMBus(busID)
	if e != nil {
		t.Logf("Error opening SMBus %d: %s\n", busID, e)
		t.Logf("NOTE: This test expects to be running on a system where "+
			"/dev/i2c-%d exists and is valid.\n", busID)
		t.FailNow()
	}
	t.Logf("Opened SMBus %d OK. Funcs = 0x%08x\n", busID, bus.Funcs)
	e = bus.Close()
	if e != nil {
		t.Logf("Error closing SMBus object: %s\n", e)
		t.Fail()
	}
}
