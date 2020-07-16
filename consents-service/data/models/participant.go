package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"time"
)

type Participant struct {
	ID              uuid.UUID       `json:"id" db:"id"`
	DefaultConsent  DefaultConsent  `json:"default_consent" has_one:"default_consent"`
	ProjectConsents ProjectConsents `json:"project_consents" has_many:"project_consents"`
}

// String is not required by pop and may be deleted
func (s Participant) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Participants is not required by pop and may be deleted
type Participants []Participant

// String is not required by pop and may be deleted
func (s Participants) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Participant) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Participant) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Participant) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
