package paunch

import (
	"errors"
	"fmt"
	gl "github.com/chsc/gogl/gl21"
	"io/ioutil"
	"strings"
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
func (effect *Effect) SetVariablei(variable string, val int) error {

	var currEffect gl.Int
	gl.GetIntegerv(gl.CURRENT_PROGRAM, &currEffect)
	if gl.Uint(currEffect) != effect.program {
		return errors.New("effect is not currently in use")
	}

	effect.checkUniformVariable(variable)

	gl.Uniform1i(effect.uniforms[variable], gl.Int(val))
	return nil
}

// SetVariable2i sets a specified variable to the two supplied integers to be
// passed into an effect.
func (effect *Effect) SetVariable2i(variable string, val1 int, val2 int) error {

	var currEffect gl.Int
	gl.GetIntegerv(gl.CURRENT_PROGRAM, &currEffect)
	if gl.Uint(currEffect) != effect.program {
		return errors.New("effect is not currently in use")
	}

	effect.checkUniformVariable(variable)

	gl.Uniform2i(effect.uniforms[variable], gl.Int(val1), gl.Int(val2))
	return nil
}

// SetVariable3i sets a specified variable to the three supplied integers to be
// passed into an effect.
func (effect *Effect) SetVariable3i(variable string, val1 int, val2 int, val3 int) error {

	var currEffect gl.Int
	gl.GetIntegerv(gl.CURRENT_PROGRAM, &currEffect)
	if gl.Uint(currEffect) != effect.program {
		return errors.New("effect is not currently in use")
	}

	effect.checkUniformVariable(variable)

	gl.Uniform3i(effect.uniforms[variable], gl.Int(val1), gl.Int(val2), gl.Int(val3))
	return nil
}

// SetVariable4i sets a specified variable to the four supplied integers to be
// passed into an effect.
func (effect *Effect) SetVariable4i(variable string, val1 int, val2 int, val3 int, val4 int) error {

	var currEffect gl.Int
	gl.GetIntegerv(gl.CURRENT_PROGRAM, &currEffect)
	if gl.Uint(currEffect) != effect.program {
		return errors.New("effect is not currently in use")
	}

	effect.checkUniformVariable(variable)

	gl.Uniform4i(effect.uniforms[variable], gl.Int(val1), gl.Int(val2), gl.Int(val3), gl.Int(val4))
	return nil
}

// SetVariablef sets a specified variable to the supplied integer to be passed
// into an effect.
func (effect *Effect) SetVariablef(variable string, val float32) error {

	var currEffect gl.Int
	gl.GetIntegerv(gl.CURRENT_PROGRAM, &currEffect)
	if gl.Uint(currEffect) != effect.program {
		return errors.New("effect is not currently in use")
	}

	effect.checkUniformVariable(variable)

	gl.Uniform1f(effect.uniforms[variable], gl.Float(val))
	return nil
}

// SetVariable2f sets a specified variable to the two supplied integers to
// be passed into an effect.
func (effect *Effect) SetVariable2f(variable string, val1 float32, val2 float32) error {

	var currEffect gl.Int
	gl.GetIntegerv(gl.CURRENT_PROGRAM, &currEffect)
	if gl.Uint(currEffect) != effect.program {
		return errors.New("effect is not currently in use")
	}

	effect.checkUniformVariable(variable)

	gl.Uniform2f(effect.uniforms[variable], gl.Float(val1), gl.Float(val2))
	return nil
}

// SetVariable3f sets a specified variable to the three supplied integers to
// be passed into an effect.
func (effect *Effect) SetVariable3f(variable string, val1 float32, val2 float32, val3 float32) error {

	var currEffect gl.Int
	gl.GetIntegerv(gl.CURRENT_PROGRAM, &currEffect)
	if gl.Uint(currEffect) != effect.program {
		return errors.New("effect is not currently in use")
	}

	effect.checkUniformVariable(variable)

	gl.Uniform3f(effect.uniforms[variable], gl.Float(val1), gl.Float(val2), gl.Float(val3))
	return nil
}

// SetVariable4f sets a specified variable to the four supplied integers to
// be passed into an effect.
func (effect *Effect) SetVariable4f(variable string, val1 float32, val2 float32, val3 float32, val4 float32) error {

	var currEffect gl.Int
	gl.GetIntegerv(gl.CURRENT_PROGRAM, &currEffect)
	if gl.Uint(currEffect) != effect.program {
		return errors.New("effect is not currently in use")
	}

	effect.checkUniformVariable(variable)

	gl.Uniform4f(effect.uniforms[variable], gl.Float(val1), gl.Float(val2), gl.Float(val3), gl.Float(val4))
	return nil
}

func compileShader(shaderType ShaderType, scripts []*gl.Char) (gl.Uint, error) {

	shaderID := gl.CreateShader(gl.Enum(shaderType))
	gl.ShaderSource(shaderID, gl.Sizei(len(scripts)), &scripts[0], nil)
	gl.CompileShader(shaderID)
	if err := checkShaderCompiled(shaderID, shaderType); err != nil {
		return 0, err
	}

	return shaderID, checkForErrors()
}

func checkShaderCompiled(id gl.Uint, shaderType ShaderType) error {

	var shader string
	if shaderType == Vertex {
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

func checkIfShaderFile(name string) ShaderType {

	split := strings.Split(name, ".")

	if split[len(split)-1] == "vert" {
		return Vertex
	} else if split[len(split)-1] == "frag" {
		return Fragment
	}

	return notShader
}

// NewEffectFromDirectory creates a new Effect object based on the shader
// directory given.
func NewEffectFromDirectory(directory string) (*Effect, error) {

	effect := &Effect{}

	effect.uniforms = make(map[string]gl.Int)

	programID := gl.CreateProgram()

	var vscript, fscript []*gl.Char

	files, ioErr := ioutil.ReadDir(directory)
	if ioErr != nil {
		return effect, ioErr
	}
	for _, val := range files {
		name := val.Name()
		if shaderType := checkIfShaderFile(name); !val.IsDir() && shaderType != notShader {
			text, err := loadTextFile(directory + name)
			if err != nil {
				return effect, err
			}
			if shaderType == Vertex {
				vscript = append(vscript, text)
			} else if shaderType == Fragment {
				fscript = append(fscript, text)
			}
		}
	}

	vshaderID, vertErr := compileShader(Vertex, vscript)
	if vertErr != nil {
		return effect, vertErr
	}
	gl.AttachShader(programID, vshaderID)
	fshaderID, fragErr := compileShader(Fragment, fscript)
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

// NewEffectFromFiles creates a new Effect object based on the shader file
// names given.
func NewEffectFromFiles(filenames []string) (*Effect, error) {

	effect := &Effect{}

	effect.uniforms = make(map[string]gl.Int)

	programID := gl.CreateProgram()

	var vscript, fscript []*gl.Char

	for _, val := range filenames {
		if shaderType := checkIfShaderFile(val); shaderType != notShader {
			text, err := loadTextFile(val)
			if err != nil {
				return effect, err
			}
			if shaderType == Vertex {
				vscript = append(vscript, text)
			} else if shaderType == Fragment {
				fscript = append(fscript, text)
			}
		}
	}

	vshaderID, vertErr := compileShader(Vertex, vscript)
	if vertErr != nil {
		return effect, vertErr
	}
	gl.AttachShader(programID, vshaderID)
	fshaderID, fragErr := compileShader(Fragment, fscript)
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

// NewEffectFromStrings creates a new Effect object based on the given strings
// containing GLSL shader script. The types values should correlate with the
// type of shader the script is describing.
func NewEffectFromStrings(text []string, types []ShaderType) (*Effect, error) {

	effect := &Effect{}

	effect.uniforms = make(map[string]gl.Int)

	programID := gl.CreateProgram()

	var vscript, fscript []*gl.Char

	for i, val := range text {
		if types[i] == Vertex {
			vscript = append(vscript, gl.GLString(val))
		} else if types[i] == Fragment {
			fscript = append(fscript, gl.GLString(val))
		}
	}

	vshaderID, vertErr := compileShader(Vertex, vscript)
	if vertErr != nil {
		return effect, vertErr
	}
	gl.AttachShader(programID, vshaderID)
	fshaderID, fragErr := compileShader(Fragment, fscript)
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
