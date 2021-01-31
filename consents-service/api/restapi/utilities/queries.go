package utilities

import (
	"fmt"

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
// FindOneDefaultConsent returns the DefaultConsent in the database corresponding to the given Participant ID (or nil if no match is found)
func FindOneDefaultConsent(participantID string, tx *pop.Connection) (*datamodels.DefaultConsent, error) {
	defaultConsent := &datamodels.DefaultConsent{}
	err := tx.Where("participant_id in (?)", participantID).First(defaultConsent)
	if err != nil {
		return nil, err
	}
	return defaultConsent, nil
}

// FindOneProjectConsent returns the ProjectConsent in the database corresponding to the given Participant and ProjectApplication IDs
// (or nil if no match is found)
func FindOneProjectConsent(participantID string, projectApplicationID int32, tx *pop.Connection) (*datamodels.ProjectConsent, error) {
	projectConsent := &datamodels.ProjectConsent{}
	query := fmt.Sprintf("participant_id = '%s' AND project_application_id = '%d'",
		participantID, projectApplicationID)
	err := tx.Where(query).First(projectConsent)
	if err != nil {
		return nil, err
	}
	return projectConsent, nil
}
