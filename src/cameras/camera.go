package cameras

import (
	"util"
)

type Camera interface {
	MakeWorldSpaceRay(i, j int, sample [2]float32) *util.Ray
}
