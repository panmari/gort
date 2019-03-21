package materials

import (
	"github.com/barnex/fmath"
	"github.com/panmari/gort/util"
	"github.com/ungerik/go3d/vec3"
)

type grid struct {
	tileMaterial, lineMaterial util.Material
	tileSize, lineSize         float32
}

func (m *grid) EvaluateEmission(hit *util.Hitrecord, wOut *vec3.T) vec3.T {
	return vec3.T{}
}
func (m *grid) GetEmissionSample(hit *util.Hitrecord, sample [2]float32) *vec3.T {
	return nil
}

func (m *grid) GetShadingSample(hit *util.Hitrecord, sample [2]float32) *vec3.T {
	return nil
}

func (m *grid) DoesCastShadows() bool {
	return true
}

func (m *grid) EvaluateSpecularReflection(hit *util.Hitrecord) (util.ShadingSample, bool) {
	return util.ShadingSample{}, false
}

func (m *grid) EvaluateBRDF(hit *util.Hitrecord, wOut, wIn *vec3.T) vec3.T {
	for _, d := range hit.Position {
		moddedD := fmath.Mod(d, m.tileSize+m.lineSize)
		if moddedD < 0 && moddedD > -m.lineSize || moddedD > m.tileSize {
			// On line, return immediately.
			return m.lineMaterial.EvaluateBRDF(hit, wOut, wIn)
		}
	}
	return m.tileMaterial.EvaluateBRDF(hit, wOut, wIn)
}

// NewGrid returns a new material that forms a grid of tileMaterial and lineMaterial (with their respective size).
// The first tile is at [(0, 0, 0), (tileSize, tileSize, tileSize)] and surrounded by a line of lineSize in every direction.
// TODO(panmari): Does currently not work with materials that use other methods than EvaluateBRDF.
func NewGrid(tileMaterial, lineMaterial util.Material, tileSize, lineSize float32) util.Material {
	return &grid{
		tileMaterial: tileMaterial,
		lineMaterial: lineMaterial,
		tileSize:     tileSize,
		lineSize:     lineSize,
	}
}
