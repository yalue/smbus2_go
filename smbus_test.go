package smbus_go

import (
	"testing"
)

// We don't use cgo for the main library, but we use it for testing; need to
// ensure that the struct sizes match what's expected by the C library.

func TestOpenBus(t *testing.T) {
	busID := 0
	bus, e := NewSMBus(busID)
	if e != nil {
		t.Logf("Error opening SMBus %d: %s\n", busID, e)
		t.Logf("NOTE: This test expects to be running on a system where "+
			"/dev/i2c-%d exists and is valid.\n", busID)
		t.FailNow()
	}
	t.Logf("Opened SMBus %d OK. Funcs = %s\n", busID, bus.Funcs)
	funcsList := bus.Funcs.GetStringsList()
	for i, s := range funcsList {
		t.Logf("  Function %d: %s\n", i, s)
	}
	e = bus.Close()
	if e != nil {
		t.Logf("Error closing SMBus object: %s\n", e)
		t.Fail()
	}
}
