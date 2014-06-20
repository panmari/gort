package lights

import (
	"util"
)

type LightGeometry interface {
	Sample(sample [2]float32) *util.Hitrecord
}
