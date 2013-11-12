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
	shaders  map[string]gl.Uint
	programs map[string]gl.Uint
}

func loadTextFile(filename string) ([]byte, error) {

	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	return text, nil
}

func (effect *Effect) Init() error {

	effect.shaders = make(map[string]gl.Uint)
	effect.programs = make(map[string]gl.Uint)

	return nil
}

// .NewEffect adds a new effect to the Effect object from a GLSL shader file.
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

// .NewEffectList adds a new effect list to the Effect object. Effect lists are
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

// .UseEffectList activates an effect list to be used on the following frames.
func (effect *Effect) UseEffectList(name string) error {

	if _, ok := effect.programs[name]; !ok {
		return errors.New(fmt.Sprintf("Effect list %s does not exist", name))
	}

	gl.UseProgram(effect.programs[name])

	return checkForErrors()
}
