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
