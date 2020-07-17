package handlers

import (
	apimodels "github.com/dycons/consents/model-vs/api/models"
	datamodels "github.com/dycons/consents/consents-service/data/models" // TODO rm
	"github.com/dycons/consents/model-vs/api/restapi/operations"
	"github.com/dycons/consents/consents-service/transformers" // TODO rm
	"github.com/dycons/consents/model-vs/api/restapi/utilities"
	"github.com/dycons/consents/model-vs/errors"
	"github.com/dycons/consents/utilities/log"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
)

// PostParticipant processes a Participant+DefaultConsent resource posted by the API request and creates it into the database.
// It then returns the URL location of this Participant, along with its uuid (the {study_identifier} parameter in the API).
func PostParticipant(params operations.PostParticipantParams) middleware.Responder {
	tx, errPayload := utilities.ConnectDevelopment(params.HTTPRequest)
	if errPayload != nil {
		return operations.NewPostParticipantServerError().WithPayload(errPayload)
	}

	// Transform DefaultConsent model API -> data
	newDefaultConsent, errPayload := defaultConsentAPIToDataModel(*params.DefaultConsent, params.HTTPRequest, tx)
	if errPayload != {
		return operations.NewPostParticipantInternalServerError().WithPayload(errPayload)
	}
	*newDefaultConsent.Participant := *Participant{}

	// Eager create DefaultConsent+Participant in DB. Only proceed if creation succeeds.
	err := tx.Eager().Create(newDefaultConsent)
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Creating into database failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPostParticipantInternalServerError().WithPayload(errPayload)
	}

	// Return Participant (with only the StudyIdentifier property populated) and Location URL header
	createdParticipant, errPayload := participantDataToAPIModel(datamodels.Participant{ID: *newDefaultConsent.ParticipantID})
	if errPayload != {
		return operations.NewPostParticipantInternalServerError().WithPayload(errPayload)
	}
	location := params.HTTPRequest.URL.Host + params.HTTPRequest.URL.EscapedPath() +
		"/" + createdParticipant.StudyIdentifier.String() + "/default_consents"
	return operations.NewPostParticipantCreated().WithPayload(createdParticipant).WithLocation(location)
}

// TODO remove here, move to generics package

// defaultConsentAPIToDataModel transforms an api.models representation of the DefaultConsent from the Go-Swagger-
// defined API to a data.models representation of the DefaultConsent from the pop ORM.
// This allows for the movement of DefaultConsent data from the server to the database for POST/PUT/DELETE
// requests.
// The transformed DefaultConsent is validated within this function prior to its return.
// An *apimodels.Error pointer is returned alongside the transformed DefaultConsent for ease of error
// response, as it can be used as the response payload immediately.
func defaultConsentAPIToDataModel(apiDefaultConsent apimodels.DefaultConsent, HTTPRequest *http.Request, tx *pop.Connection) (*datamodels.DefaultConsent, *apimodels.Error) {
	dataDefaultConsent, err := transformers.defaultConsentAPIToData(apiDefaultConsent)
	if err != nil {
		log.Write(HTTPRequest, 500000, err).Error("Failed transformation of DefaultConsent from api to data model")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	validationErrors, err := dataDefaultConsent.Validate(tx)
	if err != nil {
		log.Write(HTTPRequest, 500000, err).Error("Data schema validation for data-model DefaultConsent failed upon transformation with the following validation errors:\n" +
			validationErrors.Error())
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	return dataDefaultConsent, nil
}

// TODO remove here, more to generics package

// participantDataToAPIModel transforms a data.models representation of the Participant from the pop ORM-like
// to an api.models representation of the Participant from the Go-Swagger-defined API.
// This allows for the movement of Participant data from the database to the server for GET requests.
// An *apimodels.Error pointer is returned alongside the transformed Participant for ease of error
// response, as it can be used as the response payload immediately.
func participantDataToAPIModel(dataParticipant datamodels.Participant, HTTPRequest *http.Request) (*apimodels.Participant, *apimodels.Error) {
	apiParticipant, err := transformers.ParticipantDataToAPI(dataParticipant)
	if err != nil {
		log.Write(HTTPRequest, 500000, err).Error("Failed transformation of Participant from data to api model")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	err = apiParticipant.Validate(strfmt.NewFormats())
	if err != nil {
		log.Write(HTTPRequest, 500000, err).Error("API schema validation for API-model Participant failed upon transformation")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	return apiParticipant, nil
}
