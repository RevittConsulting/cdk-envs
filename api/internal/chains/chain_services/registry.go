package chain_services

const (
	Block = "block"
	Logs  = "logs"
)

type IService interface {
	Start() error
	Stop() error
}

type Registry struct {
	services map[string]IService
}

func NewRegistry() *Registry {
	return &Registry{
		services: make(map[string]IService),
	}
}

func (r *Registry) Register(name string, service IService) {
	r.services[name] = service
}

func (r *Registry) StartAll() error {
	for _, service := range r.services {
		if err := service.Start(); err != nil {
			return err
		}
	}
	return nil
}

func (r *Registry) StartServices(name []string) error {
	for _, serviceName := range name {
		service := r.services[serviceName]
		if err := service.Start(); err != nil {
			return err
		}
	}
	return nil
}

func (r *Registry) StopAll() error {
	for _, service := range r.services {
		if err := service.Stop(); err != nil {
			return err
		}
	}
	return nil
}

func (r *Registry) GetService(name string) IService {
	return r.services[name]
}
