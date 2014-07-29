// Package paunch is a 2D game engine written in Go.
package paunch

func init() {

	paunchWindow = newWindow()
}

// Start starts the Paunch system by opening the window and beginning to
// recieve input. Some Paunch operations may fail if this function has not yet
// been called, including all drawing commands. The version variable sets how
// Paunch initializes OpenGL. VersionOld initializes OpenGL 2.1. A few
// Paunch functionalities are not avaliable in this version. VersionNew
// initializes OpenGL 3.0, where all functionalities are avaliable.
// VersionAutomatic initializes the highest OpenGL version that the graphics
// chip will support.
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
