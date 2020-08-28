package handlers

import (
	// TODO rm

	"github.com/go-openapi/runtime/middleware" // TODO rm
	"github.com/gobuffalo/pop"

	apimodels "github.com/dycons/consents/consents-service/api/models"
	"github.com/dycons/consents/consents-service/api/restapi/operations"
	"github.com/dycons/consents/consents-service/api/restapi/utilities" // TODO rm
	"github.com/dycons/consents/consents-service/errors"
	"github.com/dycons/consents/consents-service/transformers" // TODO rm
	"github.com/dycons/consents/consents-service/utilities/log"
)

// GetOneDefaultConsent fetches the DefaultConsent for the requested defaultConsent.
func GetOneDefaultConsent(params operations.GetOneDefaultConsentParams, tx *pop.Connection) middleware.Responder {
	// Transform the Participant UUID from the api format to the data format
	participantID, err := transformers.UUIDAPIToData(params.StudyIdentifier, "StudyIdentifier")
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Transforming the StudyIdentifier from API to data formats failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewInitializeProjectConsentInternalServerError().WithPayload(errPayload)
	}

	// Find the DefaultConsent associated with the uuid given in the request
	dataDefaultConsent, err := utilities.GetDefaultConsentByParticipantID(participantID.String(), tx)
	if err != nil {
		message := "This DefaultConsent cannot be found."
		var code int64 = 404001

		log.Write(params.HTTPRequest, code, err).Warn(message)
		errPayload := &apimodels.Error{Code: &code, Message: &message}
		return operations.NewGetOneDefaultConsentNotFound().WithPayload(errPayload)
	}

	// Transform the resource from data model to api model to prepare for return
	apiDefaultConsent, errPayload := defaultConsentDataToAPIModel(*dataDefaultConsent, params.HTTPRequest)
	if errPayload != nil {
		return operations.NewGetOneDefaultConsentInternalServerError().WithPayload(errPayload)
	}

	return operations.NewGetOneDefaultConsentOK().WithPayload(apiDefaultConsent)
}
