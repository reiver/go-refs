package refsdriver


var (
	Registry Registrar = PromoteBasicRegistrar( newBasicRegistrar() )
)


type internalRegistrar map[string]Driver


func newBasicRegistrar() BasicRegistrar {
	var registrar internalRegistrar = make(map[string]Driver)

	return registrar
}


func (registrar internalRegistrar) Register(name string, driver Driver) error {

	registrar[name] = driver

	return nil
}

func (registrar internalRegistrar) Fetch(name string) (Driver, error) {

	driver, ok := registrar[name]
	if !ok {
		return nil, errNotFoundComplainer
	}

	return driver, nil
}
