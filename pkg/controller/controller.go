package controller

import (
	"github.com/golang/glog"
)

type ServiceInfo struct {
	Name string
}

// Service is the interface that each service must satisfy in order to be able
// to run.
type Service interface {
	Start() error
	Stop() error
	Reload() error
	Info() (*ServiceInfo, error)
}

// Controller controls the service instantiating, terminating and reloading.
type Controller struct {
	poolTimeSeconds int
	servicePool     []Service
}

// Start tries to start a new service..
func (c *Controller) Start() error {
	for _, service := range c.servicePool {
		err := service.Start()
		// Don't panic, just log the error from the service.
		// NOTE: errors need to be detailed to streamline debugging.
		if err != nil {
			info, err := service.Info()
			if err != nil {
				return err
			}
			glog.Warningf("Failed to start service %s", info.Name)
		}
	}
	return nil
}

// Stop stops a currently running service.
func (c *Controller) Stop() error {
	return nil
}

// Reload reloads the services
func (c *Controller) Reload() error {
	return nil
}

// LoadServices validates the services that are going to be run and adds them to
// the controllers service pool.
func (c *Controller) LoadServices() error {
	return nil
}

// newController creates a new controller.
func newController(poolTime int, services []Service) *Controller {
	return &Controller{
		poolTimeSeconds: poolTime,
	}
}
