package models

import (
	"encoding/json"
	customValidators "github.com/dycons/consents/utilities/validators"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
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
	GeneticConsentStyle  int          `json:"genetic_consent_style" db:"genetic_consent_style"`
	ClinicalConsentStyle int          `json:"clinical_consent_style" db:"clinical_consent_style"`
}

// These constants capture the enum values for the genetic & clinical "Consent Styles",
// 		the styles of default consent that may be used to initialize participation in new secondary-use projects.
const (
	start int = iota // start enum, only used in validation

	// SecondaryUseForbidden :	Do not ever share associated data with new secondary projects.
	SecondaryUseForbidden

	// OptIn :	By default, do not share associated data with new secondary projects.
	// 			Participant will opt-in to the secondary-use projects they desire to participate in.
	OptIn

	// OptOut :	By default, share associated data with new secondary projects.
	// 			Participant will opt-out of the secondary-use projects that they do not desire to participate in.
	OptOut

	end // end enum, only used in validation
)

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
	return validate.Validate(
		&validators.UUIDIsPresent{Field: d.ParticipantID, Name: "ParticipantID"},
		&customValidators.EnumIsInRange{Field: d.GeneticConsentStyle, Name: "GeneticConsentStyle", Start: start, End: end},
		&customValidators.EnumIsInRange{Field: d.ClinicalConsentStyle, Name: "ClinicalConsentStyle", Start: start, End: end},
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
