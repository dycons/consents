/*
Package utilities implements general-purpose utility functions for use by the restapi handlers.
*/
package utilities

import (
	datamodels "github.com/dycons/consents/consents-service/data/models"
	"github.com/gobuffalo/pop"
)

// addAND only adds an AND to the given conditions string if it already has contents.
func addAND(conditions string) string {
	if conditions == "" {
		return ""
	}
	return conditions + " AND "
}

// TODO remove, make generic in consents-service/api/generics/generic_resource_utilities.go
// GetParticipantByID returns the Participant in the database corresponding to the given ID (or nil if no match is found)
func GetParticipantByID(id string, tx *pop.Connection) (*datamodels.Participant, error) {
	Participant := &datamodels.Participant{}
	err := tx.Find(Participant, id)
	return Participant, err
}
