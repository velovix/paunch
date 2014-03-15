package paunch

import (
	"errors"
	"fmt"
	gl "github.com/chsc/gogl/gl33"
	"io/ioutil"
	"strings"
)

const (
	vertex    = gl.VERTEX_SHADER
	fragment  = gl.FRAGMENT_SHADER
	notShader = -1
)

// Effect is an object that manages shaders.
type Effect struct {
	uniforms map[string]gl.Int
	program  gl.Uint
}

func loadTextFile(filename string) (*gl.Char, error) {

	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return gl.GLString(""), err
	}

	return gl.GLString(string(text)), nil
}

func (effect *Effect) checkUniformVariable(variable string) {

	if _, ok := effect.uniforms[variable]; !ok {
		effect.uniforms[variable] = gl.GetUniformLocation(effect.program, gl.GLString(variable))
	}
}

// SetVariablei sets a specified variable to the supplied integer to be passed
// into an effect.
func (effect *Effect) SetVariablei(variable string, val int) {

	effect.checkUniformVariable(variable)

	gl.Uniform1i(effect.uniforms[variable], gl.Int(val))
}

// SetVariable2i sets a specified variable to the two supplied integers to be
// passed into an effect.
func (effect *Effect) SetVariable2i(variable string, val1 int, val2 int) {

	effect.checkUniformVariable(variable)

	gl.Uniform2i(effect.uniforms[variable], gl.Int(val1), gl.Int(val2))
}

// SetVariable3i sets a specified variable to the three supplied integers to be
// passed into an effect.
func (effect *Effect) SetVariable3i(variable string, val1 int, val2 int, val3 int) {

	effect.checkUniformVariable(variable)

	gl.Uniform3i(effect.uniforms[variable], gl.Int(val1), gl.Int(val2), gl.Int(val3))
}

// SetVariable4i sets a specified variable to the four supplied integers to be
// passed into an effect.
func (effect *Effect) SetVariable4i(variable string, val1 int, val2 int, val3 int, val4 int) {

	effect.checkUniformVariable(variable)

	gl.Uniform4i(effect.uniforms[variable], gl.Int(val1), gl.Int(val2), gl.Int(val3), gl.Int(val4))
}

// SetVariableui sets a specified variable to the supplied integer to be passed
// into an effect.
func (effect *Effect) SetVariableui(variable string, val uint) {

	effect.checkUniformVariable(variable)

	gl.Uniform1ui(effect.uniforms[variable], gl.Uint(val))
}

// SetVariable2ui sets a specified variable to the two supplied integers to be
// passed into an effect.
func (effect *Effect) SetVariable2ui(variable string, val1 uint, val2 uint) {

	effect.checkUniformVariable(variable)

	gl.Uniform2ui(effect.uniforms[variable], gl.Uint(val1), gl.Uint(val2))
}

// SetVariable3ui sets a specified variable to the three supplied integers to
// be passed into an effect.
func (effect *Effect) SetVariable3ui(variable string, val1 uint, val2 uint, val3 uint) {

	effect.checkUniformVariable(variable)

	gl.Uniform3ui(effect.uniforms[variable], gl.Uint(val1), gl.Uint(val2), gl.Uint(val3))
}

// SetVariable4ui sets a specified variable to the four supplied integers to
// be passed into an effect.
func (effect *Effect) SetVariable4ui(variable string, val1 uint, val2 uint, val3 uint, val4 uint) {

	effect.checkUniformVariable(variable)

	gl.Uniform4ui(effect.uniforms[variable], gl.Uint(val1), gl.Uint(val2), gl.Uint(val3), gl.Uint(val4))
}

// SetVariablef sets a specified variable to the supplied integer to be passed
// into an effect.
func (effect *Effect) SetVariablef(variable string, val float32) {

	effect.checkUniformVariable(variable)

	gl.Uniform1f(effect.uniforms[variable], gl.Float(val))
}

// SetVariable2f sets a specified variable to the two supplied integers to
// be passed into an effect.
func (effect *Effect) SetVariable2f(variable string, val1 float32, val2 float32) {

	effect.checkUniformVariable(variable)

	gl.Uniform2f(effect.uniforms[variable], gl.Float(val1), gl.Float(val2))
}

// SetVariable3f sets a specified variable to the three supplied integers to
// be passed into an effect.
func (effect *Effect) SetVariable3f(variable string, val1 float32, val2 float32, val3 float32) {

	effect.checkUniformVariable(variable)

	gl.Uniform3f(effect.uniforms[variable], gl.Float(val1), gl.Float(val2), gl.Float(val3))
}

// SetVariable4f sets a specified variable to the four supplied integers to
// be passed into an effect.
func (effect *Effect) SetVariable4f(variable string, val1 float32, val2 float32, val3 float32, val4 float32) {

	effect.checkUniformVariable(variable)

	gl.Uniform4f(effect.uniforms[variable], gl.Float(val1), gl.Float(val2), gl.Float(val3), gl.Float(val4))
}

func compileShader(shaderType int, scripts []*gl.Char) (gl.Uint, error) {

	shaderID := gl.CreateShader(gl.Enum(shaderType))
	gl.ShaderSource(shaderID, gl.Sizei(len(scripts)), &scripts[0], nil)
	gl.CompileShader(shaderID)
	if err := checkShaderCompiled(shaderID, shaderType); err != nil {
		return 0, err
	}

	return shaderID, checkForErrors()
}

func checkShaderCompiled(id gl.Uint, shaderType int) error {

	var shader string
	if shaderType == vertex {
		shader = "vertex"
	} else {
		shader = "fragment"
	}
	var status gl.Int
	gl.GetShaderiv(id, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var infoLogLength gl.Int
		gl.GetShaderiv(id, gl.INFO_LOG_LENGTH, &infoLogLength)

		infoLog := make([]gl.Char, infoLogLength+1)
		gl.GetShaderInfoLog(id, gl.Sizei(infoLogLength), nil, &infoLog[0])

		return fmt.Errorf("shader %s: %s", shader, gl.GoString(&infoLog[0]))
	}

	return nil
}

func checkIfShaderFile(name string) int {

	split := strings.Split(name, ".")

	if split[len(split)-1] == "vert" {
		return vertex
	} else if split[len(split)-1] == "frag" {
		return fragment
	}

	return -1
}

// NewEffect creates a new Effect object based on the shader directory given.
func NewEffect(directory string) (Effect, error) {

	var effect Effect

	programID := gl.CreateProgram()

	var vscript, fscript []*gl.Char

	files, ioErr := ioutil.ReadDir("shaders/" + directory)
	if ioErr != nil {
		return effect, ioErr
	}
	for _, val := range files {
		name := val.Name()
		if shaderType := checkIfShaderFile(name); !val.IsDir() && shaderType != -1 {
			text, err := loadTextFile("shaders/" + directory + name)
			if err != nil {
				return effect, err
			}
			if shaderType == vertex {
				vscript = append(vscript, text)
			} else if shaderType == fragment {
				fscript = append(fscript, text)
			}
		}
	}

	vshaderID, vertErr := compileShader(vertex, vscript)
	if vertErr != nil {
		return effect, vertErr
	}
	gl.AttachShader(programID, vshaderID)
	fshaderID, fragErr := compileShader(fragment, fscript)
	if fragErr != nil {
		return effect, fragErr
	}
	gl.AttachShader(programID, fshaderID)

	gl.LinkProgram(programID)
	var status gl.Int
	gl.GetProgramiv(programID, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		return effect, errors.New("linking program")
	}

	effect.program = programID

	return effect, checkForErrors()
}
