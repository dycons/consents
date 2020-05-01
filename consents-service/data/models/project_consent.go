package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"time"
)

type ProjectConsent struct {
	ID                   uuid.UUID 			`json:"id" db:"id"`
	CreatedAt            time.Time 			`json:"created_at" db:"created_at"`
	UpdatedAt            time.Time 			`json:"updated_at" db:"updated_at"`
	StudyParticipantID   uuid.UUID 			`json:"study_participant_id" db:"study_participant_id"`
	StudyParticipant 	 *StudyParticipant 	`json:"study_participant" belongs_to:"study_participant"`
	ProjectApplicationID uuid.UUID 			`json:"project_application_id" db:"project_application_id"`
	GeneticConsent       bool      			`json:"genetic_consent" db:"genetic_consent"`
	ClinicalConsent      bool      			`json:"clinical_consent" db:"clinical_consent"`
}

// String is not required by pop and may be deleted
func (p ProjectConsent) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// ProjectConsents is not required by pop and may be deleted
type ProjectConsents []ProjectConsent

// String is not required by pop and may be deleted
func (p ProjectConsents) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *ProjectConsent) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *ProjectConsent) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *ProjectConsent) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
