package utils

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type ProgressPrinter struct {
	numWriters      int
	progressWriters []ProgressWriter
	prevProgress    []uint64
	tw              *tabwriter.Writer
	linesWritten    int
}

func (pp *ProgressPrinter) PrintProgress() {
	if pp.linesWritten > 0 {
		for i := 0; i < pp.linesWritten; i++ {
			fmt.Fprintf(pp.tw, "\033[A%c[2K", 27)
		}

		pp.linesWritten = 0
		pp.tw.Flush()
	}

	fmt.Fprintf(pp.tw, "File\tFetched\tTotal\tSpeed\tTime\n")
	pp.linesWritten += 1

	for i, pw := range pp.progressWriters {
		if !pw.startTime.IsZero() {
			percentage := int((float64(pw.fetched) / float64(pw.total)) * 100)
			fmt.Fprintf(pp.tw, "%s\t%s (%d%%)\t%s\t", pw.filename,
				humanizeBytes(pw.fetched), percentage, humanizeBytes(pw.total))

			if !pw.endTime.IsZero() {
				fmt.Fprintf(pp.tw, "\t%v", pw.endTime.Sub(pw.startTime).Round(time.Second))
			} else {
				deltaFetched := pw.fetched - pp.prevProgress[i]
				pp.prevProgress[i] = pw.fetched
				fmt.Fprintf(pp.tw, "%s/s\t%v",
					humanizeBytes(deltaFetched), time.Now().Sub(pw.startTime).Round(time.Second))
			}

			fmt.Fprintf(pp.tw, "\n")
			pp.linesWritten += 1
		}
	}

	pp.tw.Flush()
}

func NewProgressPrinter(numWriters int) ProgressPrinter {
	tw := new(tabwriter.Writer)
	tw.Init(os.Stdout, 0, 8, 4, '\t', 0)

	return ProgressPrinter{
		numWriters:      numWriters,
		progressWriters: make([]ProgressWriter, numWriters),
		prevProgress:    make([]uint64, numWriters),
		tw:              tw,
		linesWritten:    0,
	}
}
