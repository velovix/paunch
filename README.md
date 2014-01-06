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
Paunch needs a few packages before being built.

	go get github.com/chsc/gogl/gl33
	go get github.com/go-gl/glfw3

The package `glfw3` needs the C glfw3 library installed to build correctly. See
[their repository](http://github.com/go-gl/glfw3) for more information.

Installation
------------
Once you have the dependencies installed, just run go get!

	go get github.com/Velovix/paunch

Features
--------
- Shape drawing and texturing through OpenGL
- Easy-to-use effects system through GLSL shaders
- Simple window management
- Mouse and keyboard input
- Fast, optomized collision detection
- Flexibly handle complex object movement with multiple forces
- Easy event management through object polling
- Easy menu management
- More to come!

To Do
-----
- Music and sound effect support
- Joystick handling
- Support for pre-3.0 OpenGL versions
- Probably more...
