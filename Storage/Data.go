package storage

import "time"

type Result struct {
	URL          string
	StatusCode   int
	TimeDuration time.Duration
	Err error
}

type Collector struct{
	DataReport []Result
}
