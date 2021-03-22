package services

import (
	"fmt"
)

// Service Define la interface que debe implementar un servicio
type Service interface {
	Register(serviceName string)
	Deregister()
}

var (
	services map[string]Service
)

// init es la función que se ejecuta al cargar el package
func init() {
	services = make(map[string]Service)
}

// RegisterService Agrega un nuevo servicio al registry
func RegisterService(serviceName string, service Service) error {
	if _, exists := services[serviceName]; exists {
		return fmt.Errorf("service already exists: %v", serviceName)
	}
	service.Register(serviceName)
	services[serviceName] = service
	return nil
}

// GetService Devuelve el servicio 'serviceName' del registry
func GetService(serviceName string) (Service, error) {
	if service, ok := services[serviceName]; ok {
		return service, nil
	}

	return nil, fmt.Errorf("service not exists: %v", serviceName)
}

// IsRegistered Indica si el servicio ya está registrado
func IsRegistered(serviceName string) bool {
	_, ok := services[serviceName]

	return ok
}

// Deregister Elimina un servicio del registry
func Deregister(serviceName string) error {
	if service, exists := services[serviceName]; exists {
		service.Deregister()
		delete(services, serviceName)
	} else {
		return fmt.Errorf("service unregistered: %v", serviceName)
	}

	return nil
}

// DeregisterAll Elimina todos los servicios del registry
func DeregisterAll() error {

	for serviceName, service := range services {
		service.Deregister()
		delete(services, serviceName)
	}

	return nil
}
