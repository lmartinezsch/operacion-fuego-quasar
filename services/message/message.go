package message

import (
	"strings"

	"github.com/lmartinezsch/operacion-fuego-quasar/services"
)

// ServiceName: nombre sugerido para el servicio
var ServiceName string = "message"

type Service interface {
	services.Service
	GetMessage(messages ...[]string) (msg string)
}

type messageService struct {
	serviceName string
}

// NewService devuelve la implementación del servicio messageService
func NewService() Service {

	return &messageService{}
}

// Deregister Se realizan las operaciones necesarias para quitar el
// servicio del registry
func (service *messageService) Deregister() {
}

// Register Establece el nombre con el que fue registrado el servicio
func (service *messageService) Register(serviceName string) {
	service.serviceName = serviceName
}

func (service *messageService) GetMessage(messages ...[]string) (msg string) {

	if len(messages) == 0 {
		return ""
	}

	countWords := len(messages[0])
	fullMessage := make([]string, countWords)

	for _, message := range messages {
		for i, word := range message {
			if len(word) > 0 {
				fullMessage[i] = word
			}
		}
	}

	fullMessageString := strings.Join(fullMessage, " ")

	return fullMessageString
}
