package constants

// DefaultScheme is the scheme that should be prepended to user-provided
// URLs that do not specify a scheme.
const DefaultScheme string = "http"

// SupportedUrlSchemes contains the schemes that the goget application
// currently supports. It is intended to be used for checking if user-provided
// URLs are suitable.
var SupportedUrlSchemes = map[string]struct{}{
	"http":  struct{}{},
	"https": struct{}{},
}

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
)

const (
	KB = 1e3
	MB = 1e6
	GB = 1e9
	TB = 1e12
)
