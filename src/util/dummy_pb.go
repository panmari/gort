package util


type AbstractProgressBar interface {
	Finish()
	Increment() int
}

type DummyProgressBar struct {}; 

func (pb *DummyProgressBar) Finish() {};
func (pb *DummyProgressBar) Increment() int { return 0 };