package paunch

import (
	"errors"
	al "github.com/vova616/go-openal/openal"
)

var (
	paunchALDevice  *al.Device
	paunchALContext *al.Context
)

func startAudio() error {

	paunchALDevice = al.OpenDevice("")
	if paunchALDevice == nil {
		return errors.New("failed to open device")
	}

	paunchALContext = paunchALDevice.CreateContext()
	if ok := paunchALContext.Activate(); !ok {
		return errors.New("failed to make context current")
	}

	return nil
}

func stopAudio() {
	paunchALDevice.CloseDevice()
	paunchALContext.Destroy()
}
