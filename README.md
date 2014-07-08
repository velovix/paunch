Paunch
======

Paunch is a 2D game engine written in Go.

Development Status
------------------
Paunch is still heavily in development, and lacks the key features and proper
testing that characterizes a solid game engine. Please check back frequently,
as work is being done on Paunch almost every day.

Dependencies
------------
Paunch requires a few C libraries to be installed before use.

Paunch needs the GLFW3 library installed to build correctly, which you can find
on their website [here](www.glfw.org). Make sure you build GLFW3 as a _shared
library_. When using cmake, you have to include the `-DBUILD_SHARED_LIBS=on`
flag.

Paunch also requires the OpenAL library to be installed. You can find a
download [here](http://kcat.strangesoft.net/openal.html).

Installation
------------
Once you have the dependencies installed, just run go get!

	go get github.com/velovix/paunch

Notes
-----
As of right now, Paunch requires that your graphics card supports OpenGL 3.2 or
higher. This is a limitation I hope to get rid of as soon as possible.

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
- More to come!

To Do
-----
- Support for pre-3.0 OpenGL versions
- Probably more...
