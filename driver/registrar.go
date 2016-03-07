package refsdriver


// BasicRegistrar is the kernel of a Registrar.
//
// Anything that wants to implement Registrar should implement
// BasicRegistrar and then call PromoteBasicRegistrar to turn
// their BasicRegistrar into a Registrar. For example:
//
//	type myAppleBananaCherryBasicRegistrar struct {
//		//...
//	}
//	
//	func (basicRegistrar *) Fetch(string) (Driver, error) {
//		// ...
//	}
//	
//	func (basicRegistrar *) Register(string, Driver) error {
//		// ...
//	}
//	
//	func NewAppleBananaCherryRegistrar() refdriver.Registrar {
//		// ...
//		
//		basicRegistrar := myAppleBananaCherryBasicRegistrar{
//			//...
//		}
//		
//		registrar := PromoteBasicRegistrar(&basicRegistrar)
//		
//		return registrar
//	}
//
// NOTE that only the Fetch and Register methods were implemented by
// the 'myAppleBananaCherryBasicRegistrar' struct in the example (and
// MustFetch and MustRegister were NOT implemented by the
// 'myAppleBananaCherryBasicRegistrar' struct).
type BasicRegistrar interface {
	Fetch(string) (Driver, error)
	Register(string, Driver) error
}


// Registrar represents a registrar for "refs drivers".
//
// Drivers are registered by name using either the Register or
// MustRegister methods. For example:
//
//	err := myRegistrar.Register("http", httpDriver)
//
//	myRegistrar.MustRegister("ftp", ftpDriver) // This might panic()
//
//	err := myRegistrar.Register("data", dataDriver)
//
//	myRegistrar.MustRegister("javascript", javascriptDriver) // This might panic()
//
// Later, drivers are fetched by name using either the Fetch or
// MustFetch methods. For example:
//
//	driver, err := myRegistrar.Fetch("http")
//
//	driver, err := myRegistrar.Fetch("ftp")
//
//	driver := myRegistrar.MustFetch("data") // This might panic()
//
//	driver := myRegistrar.MustFetch("javascript") // This might panic()
//
// Likely you will be using a Registrar by using the refdriver.Registry
// global variable.
type Registrar interface {
	Fetch(string) (Driver, error)
	MustFetch(string) Driver
	Register(string, Driver) error
	MustRegister(string, Driver) Registrar
}


type internalPromotionRegistrar struct {
	basicRegistrar BasicRegistrar
}


// PromoteBasicRegistrar turns a BasicRegistrar into a Registrar.
//
// You would only need to use this func if you were creating your
// own Registrar. The idea is that you only implement BasicRegistrar
// and then you call PromoteBasicRegistrar to turn your BasicRegistrar
// into a full Registrar. For example:
//
//	type myAppleBananaCherryBasicRegistrar struct {
//		//...
//	}
//	
//	func (basicRegistrar *) Fetch(string) (Driver, error) {
//		// ...
//	}
//	
//	func (basicRegistrar *) Register(string, Driver) error {
//		// ...
//	}
//	
//	func NewAppleBananaCherryRegistrar() refdriver.Registrar {
//		// ...
//		
//		basicRegistrar := myAppleBananaCherryBasicRegistrar{
//			//...
//		}
//		
//		registrar := PromoteBasicRegistrar(&basicRegistrar)
//		
//		return registrar
//	}
//
// NOTE that only the Fetch and Register methods were implemented by
// the 'myAppleBananaCherryBasicRegistrar' struct in the example (and
// MustFetch and MustRegister were NOT implemented by the
// 'myAppleBananaCherryBasicRegistrar' struct).
func PromoteBasicRegistrar(basicRegistrar BasicRegistrar) Registrar {
	registrar := internalPromotionRegistrar{
		basicRegistrar:basicRegistrar,
	}

	return &registrar
}


func (registrar *internalPromotionRegistrar) Fetch(name string) (Driver, error) {
	return registrar.basicRegistrar.Fetch(name)
}


func (registrar *internalPromotionRegistrar) MustFetch(name string) Driver {
	if driver, err := registrar.Fetch(name); nil != err {
		panic(err)
	} else {
		return driver
	}
}


func (registrar *internalPromotionRegistrar) Register(name string, driver Driver) error {
	return registrar.basicRegistrar.Register(name, driver)
}


func (registrar *internalPromotionRegistrar) MustRegister(name string, driver Driver) Registrar {
	if err := registrar.Register(name, driver); nil != err {
		panic(err)
	}

	return registrar
}
