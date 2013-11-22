package paunch

import (
	"errors"
	"fmt"
	gl "github.com/chsc/gogl/gl33"
	"io/ioutil"
	"strings"
)

const (
	VERTEX   = gl.VERTEX_SHADER
	FRAGMENT = gl.FRAGMENT_SHADER
)

// Effect is an object that manages shaders.
type Effect struct {
	uniforms        map[string]gl.Int
	programs        map[string]gl.Uint
	current_program string
}

func loadTextFile(filename string) (*gl.Char, error) {

	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return gl.GLString(""), err
	}

	return gl.GLString(string(text)), nil
}

func (effect *Effect) checkUniformVariable(program string, variable string) {

	if _, ok := effect.uniforms[variable]; !ok {
		effect.uniforms[variable] = gl.GetUniformLocation(effect.programs[program], gl.GLString(variable))
	}
}

// Init() initializes the Effect object. This must be called before doing
// anything else with the object.
func (effect *Effect) Init() error {

	effect.uniforms = make(map[string]gl.Int)
	effect.programs = make(map[string]gl.Uint)

	return nil
}

// SetVariablei sets a specified variable to the supplied integer to be passed
// into an effect.
func (effect *Effect) SetVariablei(variable string, val int) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform1i(effect.uniforms[variable], gl.Int(val))
}

// SetVariable2i sets a specified variable to the two supplied integers to be
// passed into an effect.
func (effect *Effect) SetVariable2i(variable string, val1 int, val2 int) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform2i(effect.uniforms[variable], gl.Int(val1), gl.Int(val2))
}

// SetVariable3i sets a specified variable to the three supplied integers to be
// passed into an effect.
func (effect *Effect) SetVariable3i(variable string, val1 int, val2 int, val3 int) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform3i(effect.uniforms[variable], gl.Int(val1), gl.Int(val2), gl.Int(val3))
}

// SetVariable4i sets a specified variable to the four supplied integers to be
// passed into an effect.
func (effect *Effect) SetVariable4i(variable string, val1 int, val2 int, val3 int, val4 int) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform4i(effect.uniforms[variable], gl.Int(val1), gl.Int(val2), gl.Int(val3), gl.Int(val4))
}

// SetVariableui sets a specified variable to the supplied integer to be passed
// into an effect.
func (effect *Effect) SetVariableui(variable string, val uint) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform1ui(effect.uniforms[variable], gl.Uint(val))
}

// SetVariable2ui sets a specified variable to the two supplied integers to be
// passed into an effect.
func (effect *Effect) SetVariable2ui(variable string, val1 uint, val2 uint) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform2ui(effect.uniforms[variable], gl.Uint(val1), gl.Uint(val2))
}

// SetVariable3ui sets a specified variable to the three supplied integers to
// be passed into an effect.
func (effect *Effect) SetVariable3ui(variable string, val1 uint, val2 uint, val3 uint) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform3ui(effect.uniforms[variable], gl.Uint(val1), gl.Uint(val2), gl.Uint(val3))
}

// SetVariable4ui sets a specified variable to the four supplied integers to
// be passed into an effect.
func (effect *Effect) SetVariable4ui(variable string, val1 uint, val2 uint, val3 uint, val4 uint) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform4ui(effect.uniforms[variable], gl.Uint(val1), gl.Uint(val2), gl.Uint(val3), gl.Uint(val4))
}

// SetVariablef sets a specified variable to the supplied integer to be passed
// into an effect.
func (effect *Effect) SetVariablef(variable string, val float32) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform1f(effect.uniforms[variable], gl.Float(val))
}

// SetVariable2f sets a specified variable to the two supplied integers to
// be passed into an effect.
func (effect *Effect) SetVariable2f(variable string, val1 float32, val2 float32) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform2f(effect.uniforms[variable], gl.Float(val1), gl.Float(val2))
}

// SetVariable3f sets a specified variable to the three supplied integers to
// be passed into an effect.
func (effect *Effect) SetVariable3f(variable string, val1 float32, val2 float32, val3 float32) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform3f(effect.uniforms[variable], gl.Float(val1), gl.Float(val2), gl.Float(val3))
}

// SetVariable4f sets a specified variable to the four supplied integers to
// be passed into an effect.
func (effect *Effect) SetVariable4f(variable string, val1 float32, val2 float32, val3 float32, val4 float32) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform4f(effect.uniforms[variable], gl.Float(val1), gl.Float(val2), gl.Float(val3), gl.Float(val4))
}

func compileShader(shaderType int, scripts []*gl.Char) (gl.Uint, error) {

	shader_id := gl.CreateShader(gl.Enum(shaderType))
	gl.ShaderSource(shader_id, gl.Sizei(len(scripts)), &scripts[0], nil)
	gl.CompileShader(shader_id)
	if err := checkShaderCompiled(shader_id, shaderType); err != nil {
		return 0, err
	}

	return shader_id, checkForErrors()
}

func checkShaderCompiled(id gl.Uint, shaderType int) error {

	var shader string
	if shaderType == VERTEX {
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

		return errors.New(fmt.Sprintf("Shader %s: %s", shader, gl.GoString(&infoLog[0])))
	}

	return nil
}

func checkIfShaderFile(name string) int {

	split := strings.Split(name, ".")

	if split[len(split)-1] == "vert" {
		return VERTEX
	} else if split[len(split)-1] == "frag" {
		return FRAGMENT
	}

	return -1
}

// New adds a new effect to the Effect object from a folder of GLSL shader
// files.
func (effect *Effect) New(name, directory string) error {

	program_id := gl.CreateProgram()

	var vscript, fscript []*gl.Char

	files, ioErr := ioutil.ReadDir("shaders/" + directory)
	if ioErr != nil {
		return ioErr
	}
	for _, val := range files {
		name := val.Name()
		if shaderType := checkIfShaderFile(name); !val.IsDir() && shaderType != -1 {
			text, err := loadTextFile("shaders/" + directory + name)
			if err != nil {
				return err
			}
			if shaderType == VERTEX {
				vscript = append(vscript, text)
			} else if shaderType == FRAGMENT {
				fscript = append(fscript, text)
			}
		}
	}

	vshader_id, vertErr := compileShader(VERTEX, vscript)
	if vertErr != nil {
		return vertErr
	}
	gl.AttachShader(program_id, vshader_id)
	fshader_id, fragErr := compileShader(FRAGMENT, fscript)
	if fragErr != nil {
		return fragErr
	}
	gl.AttachShader(program_id, fshader_id)

	gl.LinkProgram(program_id)
	var status gl.Int
	gl.GetProgramiv(program_id, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		return errors.New("Error linking program")
	}

	effect.programs[name] = program_id

	return checkForErrors()
}

// Use activates an effect for the following draw commands.
func (effect *Effect) Use(name string) error {

	if _, ok := effect.programs[name]; !ok {
		return errors.New(fmt.Sprintf("Effect list %s does not exist", name))
	}

	gl.UseProgram(effect.programs[name])

	effect.current_program = name
	effect.uniforms = make(map[string]gl.Int)

	return checkForErrors()
}
