package handlers

import (
	"net/http" // TODO rm

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt" // TODO rm
	"github.com/gobuffalo/pop"

	apimodels "github.com/dycons/consents/consents-service/api/models"
	"github.com/dycons/consents/consents-service/api/restapi/operations"
	"github.com/dycons/consents/consents-service/api/restapi/utilities"
	datamodels "github.com/dycons/consents/consents-service/data/models" // TODO rm
	"github.com/dycons/consents/consents-service/errors"
	"github.com/dycons/consents/consents-service/transformers" // TODO rm
	"github.com/dycons/consents/consents-service/utilities/log"
)

// GetOneDefaultConsent fetches the DefaultConsent for the requested defaultConsent.
func GetOneDefaultConsent(params operations.GetOneDefaultConsentParams, tx *pop.Connection) middleware.Responder {
	// Find the DefaultConsent associated with the uuid given in the request
	dataDefaultConsent, err := utilities.GetDefaultConsentByStudyIdentifier(params.StudyIdentifier.String(), tx)
	if err != nil {
		message := "This DefaultConsent cannot be found."
		var code int64 = 404001

		log.Write(params.HTTPRequest, code, err).Warn(message)
		errPayload := &apimodels.Error{Code: &code, Message: &message}
		return operations.NewGetOneDefaultConsentNotFound().WithPayload(errPayload)
	}

	apiDefaultConsent, errPayload := defaultConsentDataToAPIModel(*dataDefaultConsent, params.HTTPRequest)
	if errPayload != nil {
		return operations.NewGetOneDefaultConsentInternalServerError().WithPayload(errPayload)
	}

	return operations.NewGetOneDefaultConsentOK().WithPayload(apiDefaultConsent)
}

// TODO remove here, more to generics package

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
