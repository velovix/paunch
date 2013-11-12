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
	shaders  []gl.Uint
	programs []gl.Uint
}

func loadTextFile(filename string) ([]byte, error) {

	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	return text, nil
}

// .NewEffect adds a new effect to the Effect object from a GLSL shader file.
func (effect *Effect) NewEffect(mode int, name string) (int, error) {

	filename := "shaders/" + name
	if mode == VERTEX {
		filename += ".vert"
	} else if mode == FRAGMENT {
		filename += ".frag"
	}
	scriptData, err := loadTextFile(filename)
	if err != nil {
		return 0, err
	}

	script := gl.GLString(string(scriptData))
	shader_id := gl.CreateShader(gl.Enum(mode))
	gl.ShaderSource(shader_id, 1, &script, nil)
	gl.CompileShader(shader_id)

	var status gl.Int
	gl.GetShaderiv(shader_id, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		return 0, errors.New(fmt.Sprintf("Shader %s: Compile error", name))
	}

	effect.shaders = append(effect.shaders, shader_id)

	return len(effect.shaders) - 1, nil
}

// .NewEffectList adds a new effect list to the Effect object. Effect lists are
// collections of effects to be applied to the renderer.
func (effect *Effect) NewEffectList(effects []int) (int, error) {

	for i, val := range effects {
		if val >= len(effect.shaders) {
			return 0, errors.New(fmt.Sprintf("Effect %d does not exist", i))
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
		return 0, errors.New("Error linking program")
	}

	effect.programs = append(effect.programs, program)

	return len(effect.programs) - 1, checkForErrors()
}

// .UseEffectList activates an effect list to be used on the following frames.
func (effect *Effect) UseEffectList(effectList int) error {

	if effectList >= len(effect.programs) {
		return errors.New(fmt.Sprintf("Effect list %d does not exist", effectList))
	}

	gl.UseProgram(effect.programs[effectList])

	return checkForErrors()
}
