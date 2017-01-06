// Package golang implements the "golang" runtime.
package golang

import (
	"github.com/apex/apex/function"
	"github.com/apex/apex/plugins/nodejs"
)

const (
	// Runtime name used by Apex
	Runtime = "golang"
)

func init() {
	function.RegisterPlugin(Runtime, &Plugin{})
}

// Plugin implementation.
type Plugin struct{}

// Open adds the shim and golang defaults.
func (p *Plugin) Open(fn *function.Function) error {
	if fn.Runtime != Runtime {
		return nil
	}

	if fn.Hooks.Build == "" {
		fn.Hooks.Build = "GOOS=linux GOARCH=amd64 go build -o main *.go"
	}

	fn.Shim = true
	fn.Runtime = nodejs.Runtime43

	if fn.Hooks.Clean == "" {
		fn.Hooks.Clean = "rm -f main"
	}

	return nil
}
