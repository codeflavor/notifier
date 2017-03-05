package service

// Info holds information about a currently running service.
type Info struct {
	Name         string
	Interval     int
	LastMessages []string
}

// Service is the interface that each service must satisfy in order to be able
// to run.
type Service interface {
	Start() (string, error)
	Stop() error
	Reload() error
	Info() (*Info, error)
}
