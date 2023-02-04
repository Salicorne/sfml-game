package graphics

import (
	"fmt"

	sfml "github.com/telroshan/go-sfml/v2/graphics"
)

type Texture struct {
	sfml_texture *sfml.Struct_SS_sfTexture
	name         string
}

type TextureManager struct {
	textures []Texture
}

func NewTextureManager() TextureManager {
	return TextureManager{
		textures: []Texture{},
	}
}

func (t *TextureManager) LoadTexture(name string, path string) error {
	txtr := sfml.SfTexture_createFromFile(path, (sfml.SfIntRect)(sfml.SwigcptrSfIntRect(0)))

	if txtr == nil || txtr.Swigcptr() == 0 {
		return fmt.Errorf("Error loading texture %s at path %s", name, path)
	}

	t.textures = append(t.textures, Texture{sfml_texture: &txtr, name: name})
	return nil
}

func (t *TextureManager) GetSfTexture(name string) (*sfml.Struct_SS_sfTexture, error) {
	for i := range t.textures {
		if t.textures[i].name == name {
			return t.textures[i].sfml_texture, nil
		}
	}
	return nil, fmt.Errorf("Texture manager has no texture named %s", name)
}

func (t *TextureManager) Cleanup() {
	for i := range t.textures {
		sfml.SfTexture_destroy(*t.textures[i].sfml_texture)
	}
}
