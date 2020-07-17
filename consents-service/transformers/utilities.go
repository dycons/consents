package transformers

import (
	"errors"
	"github.com/go-openapi/strfmt"
	"github.com/gobuffalo/uuid"
)

// uuidAPIToData safely transforms a api-model UUID to a data-model UUID.
// This never needs to be done for entity ID fields (primary keys), such as Resource.ID,
// as the primary keys are generated automatically be the ORM.
// On the other hand, foreign keys can be transformed with this function.
func uuidAPIToData(apiUUID strfmt.UUID, fieldName string) (*uuid.UUID, error) {
	dataUUID, err := uuid.FromString(apiUUID.String())
	if err != nil {
		message := "Transformation of " + fieldName + " from api to data model fails to yield valid UUID with the following errors:\n"
		return nil, errors.New(message + err.Error())
	}

	return &dataUUID, nil
}

// stringValueOrZero converts a pointer to a string to its constituent string, but handles nil pointers
// better than a simple * conversion by converting the nil values to "" (string zero value.)
func stringValueOrZero(pointer *string) string {
	if pointer == nil {
		return ""
	}

	return *pointer
}

// intValueOrZero converts a pointer to an int to its constituent int, but handles nil pointers
// better than a simple * conversion by converting the nil values to 0 (int zero value.)
func intValueOrZero(pointer *int) int {
	if pointer == nil {
		return 0
	}

	return *pointer
}

// boolValueOrZero converts a pointer to an boolean to its constituent boolean, but handles nil pointers
// better than a simple * conversion by converting the nil values to false (boolean zero value.)
func boolValueOrZero(pointer *bool) bool {
	if pointer == nil {
		return false
	}

	return *pointer
}
