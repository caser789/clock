package clock

import (
	"testing"
	"time"
)

// Ensure that the clock's After channel sends at the correct time.
func TestClock_After(t *testing.T) {
	var ok bool
	go func() {
		time.Sleep(10 * time.Millisecond)
		ok = true
	}()
	go func() {
		time.Sleep(30 * time.Millisecond)
		t.Fatal("too late")
	}()
	gosched()

	<-New().After(20 * time.Millisecond)
	if !ok {
		t.Fatal("too early")
	}
}
