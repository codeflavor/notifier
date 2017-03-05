package controller

import (
	"github.com/codeflavor/notifier/pkg/platform"
	"github.com/codeflavor/notifier/pkg/service"
	"github.com/golang/glog"
)

// Controller controls the service instantiating, terminating and reloading.
type Controller struct {
	servicePool []service.Service
}

var servicePool = []service.Service{}

// Start tries to start a new service..
func (c *Controller) Start() error {
	// this goes into a wg?
	for _, service := range c.servicePool {
		//NOTE: remember that we have a message received from the service here!
		_, err := service.Start()
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
	//for i
}

// Load validates the services that are going to be run and adds them to
// the controllers service pool.
func (c *Controller) Load() error {
	c.servicePool = servicePool
	return nil
}
