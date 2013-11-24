package storm

import (
)

// Used to structure entity retrieval queries
type FilterTriple struct {
	field string
	predicate string
	value interface{}
}

type (
	Where FilterTriple
	And Where
	Or And
	All interface{}
	Any All
)
