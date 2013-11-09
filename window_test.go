package paunch

import (
	"testing"
)

var window Window

func TestOpen(t *testing.T) {

	err := window.Open(640, 480, "Test")
	if err != nil {
		t.Errorf(".Open(640, 480, \"Test\") returned %s", err)
	}

	err = window.Destroy()
	if err != nil {
		t.Errorf(".Destroy() returned %s", err)
	}
}

func TestUpdateDisplay(t *testing.T) {

	err := window.UpdateDisplay()
	if err != nil {
		t.Errorf(".UpdateDisplay() returned %s", err)
	}
}

func TestUpdateEvents(t *testing.T) {

	err := window.UpdateEvents()
	if err != nil {
		t.Errorf(".UpdateEvents() returned %s", err)
	}
}
