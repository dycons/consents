package handlers

import (
	// TODO rm

	"github.com/go-openapi/runtime/middleware" // TODO rm
	"github.com/gobuffalo/pop"                 // TODO rm

	"github.com/dycons/consents/consents-service/api/restapi/operations"
	datamodels "github.com/dycons/consents/consents-service/data/models" // TODO rm
	"github.com/dycons/consents/consents-service/errors"                 // TODO rm
	"github.com/dycons/consents/consents-service/utilities/log"
)

// PostParticipant processes a Participant+DefaultConsent resource posted by the API request and creates it into the database.
// It then returns the URL location of this Participant, along with its uuid (the {study_identifier} parameter in the API).
func PostParticipant(params operations.PostParticipantParams, tx *pop.Connection) middleware.Responder {
	// Transform DefaultConsent model API -> data
	newDefaultConsent, errPayload := defaultConsentAPIToDataModel(*params.DefaultConsent, params.HTTPRequest, tx)
	if errPayload != nil {
		return operations.NewPostParticipantInternalServerError().WithPayload(errPayload)
	}
	newParticipant := datamodels.Participant{
		DefaultConsent: *newDefaultConsent,
	}

	// Eager create Participant+DefaultConsent in DB. Only proceed if creation succeeds.
	err := tx.Eager().Create(&newParticipant)
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Creating into database failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostParticipantInternalServerError().WithPayload(errPayload)
	}

	// Return Participant (with only the StudyIdentifier property populated) and Location URL header
	createdParticipant, errPayload := participantDataToAPIModel(newParticipant, params.HTTPRequest)
	if errPayload != nil {
		return operations.NewPostParticipantInternalServerError().WithPayload(errPayload)
	}
	location := params.HTTPRequest.URL.Host + params.HTTPRequest.URL.EscapedPath() +
		"/" + createdParticipant.StudyIdentifier.String()
	return operations.NewPostParticipantCreated().WithPayload(createdParticipant).WithLocation(location)
}
