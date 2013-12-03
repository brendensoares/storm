package storm

import (
	"errors"
	"reflect"
	"fmt"
	"github.com/brendensoares/storm/driver"
)

type Model struct {
	context interface{}
	container string
}

// Factory configures a new model instance
func Factory(context interface{}) interface{} {
	if contextPointerValue := reflect.ValueOf(context); contextPointerValue.Kind() != reflect.Ptr {
		panic("Context must be pointer")
	} else {
		// Get addressable value for context
		contextValue := contextPointerValue.Elem()
		// Make sure the context declares a Model instance
		if modelStructField, found := contextValue.Type().FieldByName("Model"); !found {
			panic("Model struct field missing")
		} else {
			// Model struct field is present
			model := contextValue.FieldByName("Model").Addr().Interface().(*Model)
			// Set context reference so model methods can access context model fields
			model.context = context
			// Configure the model's container name, default to override tag value if present
			if model.container = modelStructField.Tag.Get("container"); model.container == "" {
				// No container override found, use the context struct name as the container
				// TODO: pluralize
				model.container = contextValue.Type().Name()
			}
		}
	}
	return context
}

// Save will create or update changed fields on the model to the backend
// database.
// Save accepts either no arguments or a single map of key/value pairs
func (self *Model) Save(args ...interface{}) (id interface{}, saveError error) {
	if len(args) == 0 {
		// Use internal fields
		// Iterate all fields via reflection
		createQuery := driver.Query{}
		modelType := reflect.TypeOf(self.context).Elem()
		modelValue := reflect.ValueOf(self.context).Elem()
		for i := 0; i < modelValue.NumField(); i++ {
			fieldType := modelType.Field(i)
			fieldValue := modelValue.Field(i)
			// Ignore storm.Model field
			if fieldType.Name == "Model" {
				continue
			}
			// TODO: ignore non-data fields
			fmt.Println("DEBUG field:", fieldType, fieldValue)
			createQuery[fieldType.Name] = fieldValue.Interface()
		}
		// Create new database record
		if id, saveError = drivers[activeDriver].Create(self.container, createQuery); saveError != nil {
			fmt.Println("CREATE ERROR:", saveError)
			return
		}
	} else {
		// Save provided fields locally and in database
		// Look for single map argument
		// TODO: test for map type or return error
		var changedFields map[string]interface{}
		changedFields = args[0].(map[string]interface{})
		if changedFields == nil {
			saveError = errors.New("Invalid argument type")
			return
		}
		// Update current database record
	}
	return
}

// Get will retrieve an instance of a model or many instances based on the arguments
// passed. You can pass a number, string or filter struct
func (self *Model) Get(args ...interface{}) (result interface{}, getError error) {
	return
}

func (self *Model) Delete() (deleteError error) {
	return
}

func (self *Model) IsLoaded() (isLoaded bool) {
	return
}
