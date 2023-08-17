package bundler

import (
	"runtime"

	"github.com/michaelzx/sqlc-stand-alone-queries-struct/internal/info"
)

func projectMetadata() ([][2]string, error) {
	return [][2]string{
		{"sqlc_version", info.Version},
		{"go_version", runtime.Version()},
		{"goos", runtime.GOOS},
		{"goarch", runtime.GOARCH},
	}, nil
}
