package internal

import (
	"errors"
	"fmt"
)

type Service interface {
	Persist()
	ReadAll() (result []Comment)
}

type ServiceName string
type serviceCreateFunc func(config *Config, comment *Comment) Service

const (
	LocalStoreServiceName ServiceName = "local-store-service"
)

type ServiceFactory struct{}

var registerServiceCreateFunc = map[ServiceName]serviceCreateFunc{
	LocalStoreServiceName: newLocalStoreService,
}

func NewServiceFactory() *ServiceFactory {
	return &ServiceFactory{}
}

func (factory *ServiceFactory) CreateService(config *Config, comment *Comment) (Service, error) {
	serviceName, isExisted := comment.Labels["service"]
	if !isExisted {
		return nil, errors.New("no service specified")
	}

	createFunc, isExisted := registerServiceCreateFunc[ServiceName(serviceName)]
	if !isExisted {
		return nil, errors.New(fmt.Sprintf("can't support %s", serviceName))
	}

	return createFunc(config, comment), nil
}

func (factory *ServiceFactory) RegisterService(name ServiceName, createFunc serviceCreateFunc) error {
	if registerServiceCreateFunc == nil {
		registerServiceCreateFunc = make(map[ServiceName]serviceCreateFunc)
	}

	_, isExisted := registerServiceCreateFunc[name]
	if isExisted {
		return errors.New("also existed same service")
	}

	registerServiceCreateFunc[name] = createFunc
	return nil
}
