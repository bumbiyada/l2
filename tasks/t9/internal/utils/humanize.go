package utils

import (
	"fmt"
	"github.com/chatzikalymnios/goget/internal/constants"
)

func humanizeBytes(b uint64) string {
	var result string

	switch {
	case b >= constants.TiB:
		result = fmt.Sprintf("%.02f TiB", float64(b)/constants.TiB)
	case b >= constants.GiB:
		result = fmt.Sprintf("%.02f GiB", float64(b)/constants.GiB)
	case b >= constants.MiB:
		result = fmt.Sprintf("%.02f MiB", float64(b)/constants.MiB)
	case b >= constants.KiB:
		result = fmt.Sprintf("%.02f KiB", float64(b)/constants.KiB)
	default:
		result = fmt.Sprintf("%.02f B", float64(b))
	}

	return result
}

func humanizeBytesSI(b uint64) string {
	var result string

	switch {
	case b >= constants.TB:
		result = fmt.Sprintf("%.02f TB", float64(b)/constants.TB)
	case b >= constants.GB:
		result = fmt.Sprintf("%.02f GB", float64(b)/constants.GB)
	case b >= constants.MB:
		result = fmt.Sprintf("%.02f MB", float64(b)/constants.MB)
	case b >= constants.KB:
		result = fmt.Sprintf("%.02f KB", float64(b)/constants.KB)
	default:
		result = fmt.Sprintf("%.02f B", float64(b))
	}

	return result
}
