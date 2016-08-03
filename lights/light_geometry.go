package lights

import (
	"github.com/panmari/gort/util"
)

type LightGeometry interface {
	Sample(sample [2]float32) *util.Hitrecord
}
