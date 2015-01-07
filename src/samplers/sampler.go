package samplers

// A sampler is used to generate locations in a [0,1)x[0,1) interval.
// Every sampler is expected to have a constructor of the form
//     NewMySampler(seed int64, maxSampleCount int)
// with seed as an initializer for random location generation and
// maxSampleCount the maximum number n that will be used to call Get2DSamples.
//
// Beware, the containing arrays are only allocated once and will be reused,
// so when calling Get2DSamples again, the values from the first time
// will be overwritten.
type Sampler interface {
	Get2DSamples(n int) [][2]float32
}
