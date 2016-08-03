package cameras

import (
	"github.com/panmari/gort/util"
)

type Camera interface {
	MakeWorldSpaceRay(i, j int, sample [2]float32) *util.Ray
}
