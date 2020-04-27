package util

import "github.com/cheggaaa/pb/v3"

type AbstractProgressBar interface {
	Finish() *pb.ProgressBar
	Increment() *pb.ProgressBar
}

// DummyProgressBar does nothing.
type DummyProgressBar struct{}

func (pb *DummyProgressBar) Finish() *pb.ProgressBar    { return nil }
func (pb *DummyProgressBar) Increment() *pb.ProgressBar { return nil }
