package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"time"
)

// DefaultConsent : 	The ORM-side representation of the ConsentStyle data object.
// 						The consent to secondary use that a participant makes by default, when their data is
// 						DAC-approved for reuse. The DefaultConsent is used to initialize a ProjectConsent.
type DefaultConsent struct {
	ID                   uuid.UUID    `json:"id" db:"id"`
	UpdatedAt            time.Time    `json:"updated_at" db:"updated_at"`
	ParticipantID        uuid.UUID    `json:"participant_id" db:"participant_id"`
	Participant          *Participant `json:"participant" belongs_to:"participant"`
	GeneticConsentStyle  int          `json:"genetic_consent_style_id" db:"genetic_consent_style_id"`
	ClinicalConsentStyle int          `json:"clinical_consent_style_id" db:"clinical_consent_style_id"`
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

// TODO Compare the values to some exported? constant defining the limits of the range of enums
// TODO	Alternately, create a custom validator that compares the value to the permitted range/list of values

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (d *DefaultConsent) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: d.ParticipantID, Name: "ParticipantID"},
		&validators.IntIsGreaterThan{Field: d.GeneticConsentStyle, Name: "GeneticConsentStyle", Compared: 0},
		&validators.IntIsGreaterThan{Field: d.GeneticConsentStyle, Name: "GeneticConsentStyle", Compared: 0},
		&validators.IntIsLessThan{Field: d.ClinicalConsentStyle, Name: "ClinicalConsentStyle", Compared: 4},
		&validators.IntIsLessThan{Field: d.ClinicalConsentStyle, Name: "ClinicalConsentStyle", Compared: 4},
	), nil
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
