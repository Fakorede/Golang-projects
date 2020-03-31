// Package goversion is a package to...
package goversion

import "runtime"

// Version returns the current Go version
func Version() string {
	return runtime.Version()
}
