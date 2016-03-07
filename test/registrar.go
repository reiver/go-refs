package refstest


import (
	"github.com/reiver/go-refs/driver"
)


type internalRegistrar map[string]refsdriver.Driver


func NewRegistrar() refsdriver.Registrar {
	var registrar internalRegistrar = make(map[string]refsdriver.Driver)

	return refsdriver.PromoteBasicRegistrar(registrar)
}


func (registrar internalRegistrar) Register(name string, driver refsdriver.Driver) error {

	registrar[name] = driver

	return nil
}


func (registrar internalRegistrar) Fetch(name string) (refsdriver.Driver, error) {

	driver, ok := registrar[name]
	if !ok {
		return nil, errNotFoundComplainer
	}

	return driver, nil
}
