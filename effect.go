package paunch

import (
	"errors"
	"fmt"
	gl "github.com/chsc/gogl/gl33"
	"io/ioutil"
)

const (
	VERTEX   = gl.VERTEX_SHADER
	FRAGMENT = gl.FRAGMENT_SHADER
)

// Effect is an object that manages shaders.
type Effect struct {
	uniforms        map[string]gl.Int
	shaders         map[string]gl.Uint
	programs        map[string]gl.Uint
	current_program string
}

func loadTextFile(filename string) ([]byte, error) {

	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	return text, nil
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
	effect.shaders = make(map[string]gl.Uint)
	effect.programs = make(map[string]gl.Uint)

	return nil
}

// SetVariablei sets a specified variable to the supplied integer to be passed
// into an effects list.
func (effect *Effect) SetVariablei(variable string, val int) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform1i(effect.uniforms[variable], gl.Int(val))
}

// SetVariable2i sets a specified variable to the two supplied integers to be
// passed into an effects list.
func (effect *Effect) SetVariable2i(variable string, val1 int, val2 int) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform2i(effect.uniforms[variable], gl.Int(val1), gl.Int(val2))
}

// SetVariable3i sets a specified variable to the three supplied integers to be
// passed into an effects list.
func (effect *Effect) SetVariable3i(variable string, val1 int, val2 int, val3 int) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform3i(effect.uniforms[variable], gl.Int(val1), gl.Int(val2), gl.Int(val3))
}

// SetVariable4i sets a specified variable to the four supplied integers to be
// passed into an effects list.
func (effect *Effect) SetVariable4i(variable string, val1 int, val2 int, val3 int, val4 int) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform4i(effect.uniforms[variable], gl.Int(val1), gl.Int(val2), gl.Int(val3), gl.Int(val4))
}

// SetVariableui sets a specified variable to the supplied integer to be passed
// into an effects list.
func (effect *Effect) SetVariableui(variable string, val uint) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform1ui(effect.uniforms[variable], gl.Uint(val))
}

// SetVariable2ui sets a specified variable to the two supplied integers to be
// passed into an effects list.
func (effect *Effect) SetVariable2ui(variable string, val1 uint, val2 uint) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform2ui(effect.uniforms[variable], gl.Uint(val1), gl.Uint(val2))
}

// SetVariable3ui sets a specified variable to the three supplied integers to
// be passed into an effects list.
func (effect *Effect) SetVariable3ui(variable string, val1 uint, val2 uint, val3 uint) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform3ui(effect.uniforms[variable], gl.Uint(val1), gl.Uint(val2), gl.Uint(val3))
}

// SetVariable4ui sets a specified variable to the four supplied integers to
// be passed into an effects list.
func (effect *Effect) SetVariable4ui(variable string, val1 uint, val2 uint, val3 uint, val4 uint) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform4ui(effect.uniforms[variable], gl.Uint(val1), gl.Uint(val2), gl.Uint(val3), gl.Uint(val4))
}

// SetVariablef sets a specified variable to the supplied integer to be passed
// into an effects list.
func (effect *Effect) SetVariablef(variable string, val float32) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform1f(effect.uniforms[variable], gl.Float(val))
}

// SetVariable2f sets a specified variable to the two supplied integers to
// be passed into an effects list.
func (effect *Effect) SetVariable2f(variable string, val1 float32, val2 float32) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform2f(effect.uniforms[variable], gl.Float(val1), gl.Float(val2))
}

// SetVariable3f sets a specified variable to the three supplied integers to
// be passed into an effects list.
func (effect *Effect) SetVariable3f(variable string, val1 float32, val2 float32, val3 float32) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform3f(effect.uniforms[variable], gl.Float(val1), gl.Float(val2), gl.Float(val3))
}

// SetVariable4f sets a specified variable to the four supplied integers to
// be passed into an effects list.
func (effect *Effect) SetVariable4f(variable string, val1 float32, val2 float32, val3 float32, val4 float32) {

	effect.checkUniformVariable(effect.current_program, variable)

	gl.Uniform4f(effect.uniforms[variable], gl.Float(val1), gl.Float(val2), gl.Float(val3), gl.Float(val4))
}

// NewEffect adds a new effect to the Effect object from a GLSL shader file.
func (effect *Effect) NewEffect(mode int, name string) error {

	if _, ok := effect.shaders[name]; ok {
		return errors.New(fmt.Sprintf("Shader %s already exists", name))
	}

	filename := "shaders/" + name
	if mode == VERTEX {
		filename += ".vert"
	} else if mode == FRAGMENT {
		filename += ".frag"
	}
	scriptData, err := loadTextFile(filename)
	if err != nil {
		return err
	}

	script := gl.GLString(string(scriptData))
	shader_id := gl.CreateShader(gl.Enum(mode))
	gl.ShaderSource(shader_id, 1, &script, nil)
	gl.CompileShader(shader_id)

	var status gl.Int
	gl.GetShaderiv(shader_id, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		return errors.New(fmt.Sprintf("Shader %s: Compile error", name))
	}

	effect.shaders[name] = shader_id

	return nil
}

// NewEffectList adds a new effect list to the Effect object. Effect lists are
// collections of effects to be applied to the renderer.
func (effect *Effect) NewEffectList(name string, effects []string) error {

	for _, val := range effects {
		if _, ok := effect.shaders[val]; !ok {
			return errors.New(fmt.Sprintf("Effect %s does not exist", val))
		}
	}

	program := gl.CreateProgram()
	for _, val := range effects {
		gl.AttachShader(program, effect.shaders[val])
	}

	gl.LinkProgram(program)
	var status gl.Int
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		return errors.New("Error linking program")
	}

	effect.programs[name] = program

	return checkForErrors()
}

// UseEffectList activates an effect list to be used on the following frames.
func (effect *Effect) UseEffectList(name string) error {

	if _, ok := effect.programs[name]; !ok {
		return errors.New(fmt.Sprintf("Effect list %s does not exist", name))
	}

	gl.UseProgram(effect.programs[name])

	effect.current_program = name
	effect.uniforms = make(map[string]gl.Int)

	return checkForErrors()
}
