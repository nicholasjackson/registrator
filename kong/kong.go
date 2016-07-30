package kong

import (
	"net/url"

	"github.com/gliderlabs/registrator/bridge"
)

// Adapter implements a Registrator backend for the Kong API gateway (http://getcong.com)
type Adapter struct{}

func init() {
	f := new(Factory)
	bridge.Register(f, "kong")
}

// Factory for creating new Adapters
type Factory struct{}

// New creates a new Adapter for the given uri
func (f *Factory) New(uri *url.URL) bridge.RegistryAdapter {
	return &Adapter{}
}

// Ping tries to connect to kong and get the current status
func (k *Adapter) Ping() error {
	return nil
}

// Register registers a service with Kong
func (k *Adapter) Register(service *bridge.Service) error {
	return nil
}

// Deregister removes a service from Kong
func (k *Adapter) Deregister(service *bridge.Service) error {
	return nil
}

// Refresh does nothing for this backend
func (k *Adapter) Refresh(service *bridge.Service) error {
	return nil
}

// Services returns an array of service
func (k *Adapter) Services() ([]*bridge.Service, error) {
	return nil, nil
}
