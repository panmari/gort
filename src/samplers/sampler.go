package samplers

type Sampler interface {
	Get2DSample() ([2]float32)
}