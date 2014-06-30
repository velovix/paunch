// Package paunch is a 2D game engine written in Go. It uses OpenGL and other
// cross-platform technologies like GLFW, making games built on Paunch
// immediately portable.
package paunch

func init() {

	paunchWindow = newWindow()
}

// Start starts the Paunch system by opening the window and beginning to
// recieve input. Some Paunch operations may fail if this function has not yet
// been called, including all drawing commands.
func Start(version Version) error {

	var err error

	err = initWindows()
	if err != nil {
		return err
	}
	err = paunchWindow.open()
	if err != nil {
		return err
	}

	err = initDraw(version)
	if err != nil {
		return err
	}

	err = startAudio()
	if err != nil {
		return err
	}

	return nil
}

// Stop safely closes the Paunch instance.
func Stop() {

	paunchWindow.close()

	stopAudio()
}
