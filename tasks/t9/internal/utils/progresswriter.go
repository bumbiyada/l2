package utils

import (
	"time"
)

type ProgressWriter struct {
	filename    string
	total       uint64
	fetched     uint64
	prevFetched uint64
	startTime   time.Time
	endTime     time.Time
}

func (pw *ProgressWriter) Write(p []byte) (int, error) {
	numBytes := len(p)
	pw.fetched += uint64(numBytes)
	return numBytes, nil
}
