package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"time"
)

type DefaultConsent struct {
	ID                     uuid.UUID 			`json:"id" db:"id"`
	CreatedAt              time.Time 			`json:"created_at" db:"created_at"`
	UpdatedAt              time.Time 			`json:"updated_at" db:"updated_at"`
	StudyParticipantID     uuid.UUID 			`json:"study_participant_id" db:"study_participant_id"`
	StudyParticipant 	   *StudyParticipant 	`json:"study_participant" belongs_to:"study_participant"`
	GeneticConsentStyleID  uuid.UUID 			`json:"genetic_consent_style_id" db:"genetic_consent_style_id"`
	GeneticConsentStyle    *ConsentStyle 		`json:"genetic_consent_style" belongs_to:"consent_style"`
	ClinicalConsentStyleID uuid.UUID 			`json:"clinical_consent_style_id" db:"clinical_consent_style_id"`
	ClinicalConsentStyle   *ConsentStyle 		`json:"clinical_consent_style" belongs_to:"consent_style"`
}

// String is not required by pop and may be deleted
func (d DefaultConsent) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// DefaultConsents is not required by pop and may be deleted
type DefaultConsents []DefaultConsent

// String is not required by pop and may be deleted
func (d DefaultConsents) String() string {
	jd, _ := json.Marshal(d)
	return string(jd)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *DefaultConsent) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (d *DefaultConsent) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (d *DefaultConsent) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
