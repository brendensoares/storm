/*
	STORM is Semantically Terse Object Relational Mapping for Go
*/
package storm

import (
	"errors"
	"github.com/brendensoares/storm/driver"
)



var (
	drivers = make(map[string]driver.Driver)
	activeDriver string
)

func RegisterDriver(newDriver driver.Driver) {
	if _, duplicate := drivers[newDriver.Name()]; duplicate {
		panic("Data driver already added")
	}
	drivers[newDriver.Name()] = newDriver
}


func Connect(driverName string, driverConfig string) (connectError error) {
	if driverName == "" || driverConfig == "" {
		// Failure
		return errors.New("Invalid configuration given")
	} else {
		// Success
		// TODO: check that the driver is ready to receive requests
		if connectError = drivers[driverName].Open(driverConfig); connectError != nil {
			// Failure
			return
		} else {
			// Success
			activeDriver = driverName
		}
	}
	return
}
