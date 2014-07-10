package samplers

type Sampler interface {
	Get2DSamples(n int) [][2]float32
}
