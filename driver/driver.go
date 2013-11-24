package driver

import (
)

type Query map[string]interface{}


// DataDriver describes a common interface to establish a database connection
// with different database backends.
type Driver interface {
	// Establish a connection with the backend database
	Open(config string) (openError error)

	// Provides driver name
	Name() string

	// Provides connection configuration values
	Config() string

	// Create one new entity and return its identifier or return an error
	Create(container string, query Query) (id string, requestError error)

	// Read a single entity and return its representation or return an error
	// Must support semantic filtering
	ReadOne(container string, id interface{}) (result interface{}, requestError error)

	// Read multiple entities and return their representations or return an error
	// Must support semantic filtering
	ReadMany(container string, conditions string) (results interface{}, requestError error)

	// Update an entity's representation or return an error
	Update(container string, id interface{}, query Query) (requestError error)

	// Delete an entity or return an error
	Delete(container string, id interface{}) (requestError error)
}
