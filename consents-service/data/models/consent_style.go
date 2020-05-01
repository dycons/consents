package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"time"
)

type ConsentStyle struct {
	ID           			uuid.UUID 			`json:"id" db:"id"`
	ConsentStyle 			string    			`json:"consent_style" db:"consent_style"`
	GeneticDefaultConsents	[]DefaultConsent	`json:"genetic_default_consents" has_many:"default_consents" fk_id:"genetic_consent_style_id"`
	ClinicalDefaultConsents []DefaultConsent	`json:"clinical_default_consents" has_many:"default_consents" fk_id:"clinical_consent_style_id"`
}

// String is not required by pop and may be deleted
func (c ConsentStyle) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// ConsentStyles is not required by pop and may be deleted
type ConsentStyles []ConsentStyle

// String is not required by pop and may be deleted
func (c ConsentStyles) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *ConsentStyle) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.ConsentStyle, Name: "ConsentStyle"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *ConsentStyle) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *ConsentStyle) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
