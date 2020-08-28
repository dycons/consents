package handlers

import (
	// TODO rm

	"github.com/go-openapi/runtime/middleware" // TODO rm
	"github.com/gobuffalo/pop"

	apimodels "github.com/dycons/consents/consents-service/api/models"
	"github.com/dycons/consents/consents-service/api/restapi/operations"
	datamodels "github.com/dycons/consents/consents-service/data/models" // TODO rm
	"github.com/dycons/consents/consents-service/errors"
	"github.com/dycons/consents/consents-service/transformers"

	"github.com/dycons/consents/consents-service/utilities/log"
)

// GetProjectConsentsByParticipant fetches all of the ProjectConsents associated with a Participant
func GetProjectConsentsByParticipant(params operations.GetProjectConsentsByParticipantParams, tx *pop.Connection) middleware.Responder {
	// Transform the Participant UUID from the api format to the data format
	participantID, err := transformers.UUIDAPIToData(params.StudyIdentifier, "StudyIdentifier")
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Transforming the StudyIdentifier from API to data formats failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewGetProjectConsentsByParticipantInternalServerError().WithPayload(errPayload)
	}

	// Check that this participant exists in the database, return a 404 if they don't
	participant := datamodels.Participant{}
	participantExists, err := tx.Where("id in (?)", participantID).Exists(&participant)
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Checking for the existence of the requested Participant failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewGetProjectConsentsByParticipantInternalServerError().WithPayload(errPayload)
	}
	if !participantExists {
		message := "This Participant cannot be found."
		var code int64 = 404002

		log.Write(params.HTTPRequest, code, err).Warn(message)
		errPayload := &apimodels.Error{Code: &code, Message: &message}
		return operations.NewGetProjectConsentsByParticipantNotFound().WithPayload(errPayload)
	}

	// Find the ProjectConsents associated with the requested Participant
	dataProjectConsents := datamodels.ProjectConsents{}
	err = tx.Where("participant_id in (?)", participantID).All(&dataProjectConsents)
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Retrieving the ProjectConsents from the database failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewGetProjectConsentsByParticipantInternalServerError().WithPayload(errPayload)
	}

	// Transform the resource from data model to api model to prepare for return
	// The payload returned must be an array of pointers: []*apimodels.ProjectConsent
	var apiProjectConsents []*apimodels.ProjectConsent
	for _, dataProjectConsent := range dataProjectConsents {
		apiProjectConsent, errPayload := projectConsentDataToAPIModel(dataProjectConsent, params.HTTPRequest)
		if errPayload != nil {
			return operations.NewGetProjectConsentsByParticipantInternalServerError().WithPayload(errPayload)
		}
		apiProjectConsents = append(apiProjectConsents, apiProjectConsent)
	}

	return operations.NewGetProjectConsentsByParticipantOK().WithPayload(apiProjectConsents)
}
