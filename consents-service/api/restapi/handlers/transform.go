package handlers

import (
	"net/http"

	apimodels "github.com/dycons/consents/consents-service/api/models"
	datamodels "github.com/dycons/consents/consents-service/data/models"
	"github.com/dycons/consents/consents-service/errors"
	"github.com/dycons/consents/consents-service/transformers"
	"github.com/dycons/consents/consents-service/utilities/log"
	"github.com/go-openapi/strfmt"
	"github.com/gobuffalo/pop"
)

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

// defaultConsentAPIToDataModel transforms an api.models representation of the DefaultConsent from the Go-Swagger-
// defined API to a data.models representation of the DefaultConsent from the pop ORM.
// This allows for the movement of DefaultConsent data from the server to the database for POST/PUT/DELETE
// requests.
// The transformed DefaultConsent is validated within this function prior to its return.
// An *apimodels.Error pointer is returned alongside the transformed DefaultConsent for ease of error
// response, as it can be used as the response payload immediately.
func defaultConsentAPIToDataModel(apiDefaultConsent apimodels.DefaultConsent, HTTPRequest *http.Request, tx *pop.Connection) (*datamodels.DefaultConsent, *apimodels.Error) {
	dataDefaultConsent, err := transformers.DefaultConsentAPIToData(apiDefaultConsent)
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

// defaultConsentDataToAPIModel transforms a data.models representation of the DefaultConsent from the pop ORM-like
// to an api.models representation of the DefaultConsent from the Go-Swagger-defined API.
// This allows for the movement of DefaultConsent data from the database to the server for GET requests.
// An *apimodels.Error pointer is returned alongside the transformed DefaultConsent for ease of error
// response, as it can be used as the response payload immediately.
func defaultConsentDataToAPIModel(dataDefaultConsent datamodels.DefaultConsent, HTTPRequest *http.Request) (*apimodels.DefaultConsent, *apimodels.Error) {
	apiDefaultConsent, err := transformers.DefaultConsentDataToAPI(dataDefaultConsent)
	if err != nil {
		log.Write(HTTPRequest, 500000, err).Error("Failed transformation of DefaultConsent from data to api model")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	err = apiDefaultConsent.Validate(strfmt.NewFormats())
	if err != nil {
		log.Write(HTTPRequest, 500000, err).Error("API schema validation for API-model DefaultConsent failed upon transformation")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	return apiDefaultConsent, nil
}

// projectConsentDataToAPIModel transforms a data.models representation of the Participant from the pop ORM-like
// to an api.models representation of the Participant from the Go-Swagger-defined API.
// This allows for the movement of Participant data from the database to the server for GET requests.
// An *apimodels.Error pointer is returned alongside the transformed Participant for ease of error
// response, as it can be used as the response payload immediately.
func projectConsentDataToAPIModel(dataProjectConsent datamodels.ProjectConsent, HTTPRequest *http.Request) (*apimodels.ProjectConsent, *apimodels.Error) {
	apiProjectConsent, err := transformers.ProjectConsentDataToAPI(dataProjectConsent)
	if err != nil {
		log.Write(HTTPRequest, 500000, err).Error("Failed transformation of ProjectConsent from data to api model")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	err = apiProjectConsent.Validate(strfmt.NewFormats())
	if err != nil {
		log.Write(HTTPRequest, 500000, err).Error("API schema validation for API-model ProjectConsent failed upon transformation")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}

	return apiProjectConsent, nil
}
