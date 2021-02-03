package handlers

import (
	// TODO rm

	"net/http"

	"github.com/go-openapi/runtime/middleware" // TODO rm
	// TODO rm
	"github.com/gobuffalo/pop" // TODO rm

	apimodels "github.com/dycons/consents/consents-service/api/models"
	"github.com/dycons/consents/consents-service/api/restapi/operations"
	"github.com/dycons/consents/consents-service/api/restapi/utilities"
	datamodels "github.com/dycons/consents/consents-service/data/models" // TODO rm
	"github.com/dycons/consents/consents-service/errors"                 // TODO rm
	"github.com/dycons/consents/consents-service/transformers"
	"github.com/dycons/consents/consents-service/utilities/log"
)

// AddRemsEntitlements iterates over the array of RemsEntitlements, calling AddRemsEntitlement
// on each item.
// In this draft of the API (/v0), successfully posted items are committed in order of their appearance
// in the payload array. Once an item errors, it and all subsequent items are not committed, and
// the status for the errored item is returned.
// TODO /v1: Make the entire set of Entitlements object to a single transaction. Roll-back on all Entitlements
// in the request if a single Entitlement errors. See the following resources for hints on how to manage
// transactions with the pop orm:
// 		https://github.com/gobuffalo/buffalo-pop
// 		https://gobuffalo.io/en/docs/db/buffalo-integration/#handle-transaction-by-hand
func AddRemsEntitlements(params operations.AddRemsEntitlementsParams, tx *pop.Connection) middleware.Responder {
	var responder middleware.Responder
	var err error

	for _, RemsEntitlement := range params.RemsEntitlements {
		responder, err = AddRemsEntitlement(*RemsEntitlement, params.HTTPRequest, tx)
		if err != nil {
			break
		}
	}

	return responder
}

// AddRemsEntitlement takes the default consents for the study_identifier posted by the API request,
// and uses them to initialize a project-specific set of consents, which it creates into the database.
// In order to handle race conditions caused by the addition of new Entitlements for multiple users under
// one project Application, this handler accepts the request for processing but may fail a uniqueness
// constraint when writing the data to the database. A 202 Accepted response is returned regardless.
func AddRemsEntitlement(RemsEntitlement apimodels.RemsEntitlement, HTTPRequest *http.Request, tx *pop.Connection) (middleware.Responder, error) {
	// Transform the Participant UUID from the api format to the data format
	participantID, err := transformers.UUIDAPIToData(*RemsEntitlement.Resource, "StudyIdentifier")
	if err != nil {
		log.Write(HTTPRequest, 500000, err).Error("Transforming the resource UUID from API to data formats failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewAddRemsEntitlementsInternalServerError().WithPayload(errPayload), err
	}

	// Find the DefaultConsent associated with the uuid given in the request
	defaultConsent, err := utilities.FindOneDefaultConsent(participantID.String(), tx)
	if err != nil {
		message := "This Resource cannot be found"
		var code int64 = 404003

		log.Write(HTTPRequest, code, err).Warn(message)
		errPayload := &apimodels.Error{Code: &code, Message: &message}
		return operations.NewAddRemsEntitlementsNotFound().WithPayload(errPayload), err
	}

	// Initialize the project consents from the default settings
	geneticConsent, clinicalConsent, err := utilities.InitializeProjectConsents(*defaultConsent)
	if err != nil {
		log.Write(HTTPRequest, 500000, err).Error(err)
		errPayload := errors.DefaultInternalServerError()
		return operations.NewAddRemsEntitlementsInternalServerError().WithPayload(errPayload), err
	}

	// Create the ProjectConsent into the DB. Only proceed if creation succeeds.
	projectConsent := datamodels.ProjectConsent{
		ParticipantID:        *participantID,
		ProjectApplicationID: int(*RemsEntitlement.Application),
		GeneticConsent:       geneticConsent,
		ClinicalConsent:      clinicalConsent,
	}
	validationErrors, err := tx.ValidateAndCreate(&projectConsent)
	if validationErrors.Error() != "" { // if at least one validation error occured
		log.Write(HTTPRequest, 500000, err).Error("Data schema validation for the ProjectConsent failed with the following validation errors: " +
			validationErrors.Error())
		errPayload := errors.DefaultInternalServerError()
		return operations.NewAddRemsEntitlementsInternalServerError().WithPayload(errPayload), err
	}
	if err != nil {
		// Log the error, but continue to return the "Accepted" response instead of and Internal Error,
		// because most likely this error is a unique-index-duplication error, which means that this
		// data already exists in the database and does not need to be created into it.
		log.Write(HTTPRequest, 500000, err).Debug("Creation of the ProjectConsent into the database failed without validation errors.")
	}

	// Regardless of whether or not creation into the database succeeds, return a 202 Accepted response
	return operations.NewAddRemsEntitlementsAccepted(), nil
}
