package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"time"
)

type StudyParticipant struct {
	ID              uuid.UUID       `json:"id" db:"id"`
	CreatedAt       ime.Time        `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at" db:"updated_at"`
	DefaultConsent  DefaultConsent  `json:"default_consent" has_one:"default_consent"`
	ProjectConsents ProjectConsents `json:"project_consents" has_many:"project_consents"`
}

// String is not required by pop and may be deleted
func (s StudyParticipant) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// StudyParticipants is not required by pop and may be deleted
type StudyParticipants []StudyParticipant

// String is not required by pop and may be deleted
func (s StudyParticipants) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *StudyParticipant) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *StudyParticipant) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *StudyParticipant) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
