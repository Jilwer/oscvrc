package vrcinput

import (
	"github.com/hypebeast/go-osc/osc"
)

type localOscClient struct {
	*osc.Client
}

const (
	DefaultAddr = "127.0.0.1"
	DefaultPort = 9000
)

// Initiates an OSC client
// Default address is localhost (127.0.0.1), Default port is 9000
func NewOscClient(addr string, port int) localOscClient {
	oscClient := osc.NewClient(addr, port)
	return localOscClient{oscClient}
}
