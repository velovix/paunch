Paunch [![GoDoc](https://godoc.org/github.com/velovix/paunch?status.svg)](https://godoc.org/github.com/velovix/paunch)
======
Paunch is a 2D game engine written in Go.

Development Status
------------------
Paunch is nearing a stable release. There may still be minor API changes, but
hopefully future changes will either be bug fixes or new features.

Dependencies
------------
Paunch requires a few C libraries to be installed before use.

Paunch needs the GLFW3 library installed to build correctly, which you can find
on their website [here](www.glfw.org). Make sure you build GLFW3 as a _shared
library_. When using cmake, you have to include the `-DBUILD_SHARED_LIBS=on`
flag. You may also want to include the `-GLFW_USE_OPTIMUS_HPG=on` flag if you
are on a system with Nvidia Optimus in order to force Optimus to use the
dedicated graphics card. Intel integrated graphics sometimes have trouble
initializing legacy OpenGL functions.

Paunch also requires the OpenAL library to be installed. You can find a
download [here](http://kcat.strangesoft.net/openal.html).

Installation
------------
Once you have the dependencies installed, just run go get!

	go get github.com/velovix/paunch

Usage
-----
Please take a look at the files in the examples directory if you are looking
for help with getting started. The automatically generated documentation found
[here](https://godoc.org/github.com/velovix/paunch) is also a very helpful
reference. I will make an official guide to Paunch when I finish a stable
release, but in the mean time, feel free to email me with any questions!

Notes
-----
Your graphics card must support at least OpenGL 2.1 or you will get an error
complaining that OpenGL could not be initialized.

Features
--------
- Shape drawing and texturing through OpenGL
- Easy-to-use effects system through GLSL shaders
- Simple window management
- Mouse and keyboard input
- Joystick input
- Fast, optomized collision detection
- Flexibly handle complex object movement with multiple forces
- Easy event management through object polling
- Easy menu management
- Music and sound effect support
- Flexible OpenGL version support

To Do
-----
- ???

Known Issues
------------
- Intel HD Graphics 4000 may not initialize OpenGL correctly
