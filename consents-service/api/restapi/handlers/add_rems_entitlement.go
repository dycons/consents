package handlers

import (
	// TODO rm

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

// AddRemsEntitlement takes the default consents for the study_identifier posted by the API request,
// and uses them to initialize a project-specific set of consents, which it creates into the database.
// In order to handle race conditions caused by the addition of new Entitlements for multiple users under
// one project Application, this handler accepts the request for processing but may fail a uniqueness
// constraint when writing the data to the database. A 202 Accepted response is returned regardless.
func AddRemsEntitlement(params operations.AddRemsEntitlementParams, tx *pop.Connection) middleware.Responder {
	// Transform the Participant UUID from the api format to the data format
	participantID, err := transformers.UUIDAPIToData(*params.RemsEntitlement.Resource, "StudyIdentifier")
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Transforming the resource UUID from API to data formats failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewAddRemsEntitlementInternalServerError().WithPayload(errPayload)
	}

	// Find the DefaultConsent associated with the uuid given in the request
	defaultConsent, err := utilities.FindOneDefaultConsent(participantID.String(), tx)
	if err != nil {
		message := "This Resource cannot be found"
		var code int64 = 404003

		log.Write(params.HTTPRequest, code, err).Warn(message)
		errPayload := &apimodels.Error{Code: &code, Message: &message}
		return operations.NewAddRemsEntitlementNotFound().WithPayload(errPayload)
	}

	// Initialize the project consents from the default settings
	geneticConsent, clinicalConsent, err := utilities.InitializeProjectConsents(*defaultConsent)
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error(err)
		errPayload := errors.DefaultInternalServerError()
		return operations.NewAddRemsEntitlementInternalServerError().WithPayload(errPayload)
	}

	// Create the ProjectConsent into the DB. Only proceed if creation succeeds.
	projectConsent := datamodels.ProjectConsent{
		ParticipantID:        *participantID,
		ProjectApplicationID: int(*params.RemsEntitlement.Application),
		GeneticConsent:       geneticConsent,
		ClinicalConsent:      clinicalConsent,
	}
	validationErrors, err := tx.ValidateAndCreate(&projectConsent)
	if validationErrors.Error() != "" { // if at least one validation error occured
		log.Write(params.HTTPRequest, 500000, err).Error("Data schema validation for the ProjectConsent failed with the following validation errors: " +
			validationErrors.Error())
		errPayload := errors.DefaultInternalServerError()
		return operations.NewAddRemsEntitlementInternalServerError().WithPayload(errPayload)
	}
	if err != nil {
		// Log the error, but continue to return the "Accepted" response instead of and Internal Error,
		// because most likely this error is a unique-index-duplication error, which means that this
		// data already exists in the database and does not need to be created into it.
		log.Write(params.HTTPRequest, 500000, err).Debug("Creation of the ProjectConsent into the database failed without validation errors.")
	}

	// Regardless of whether or not creation into the database succeeds, return a 202 Accepted response
	return operations.NewAddRemsEntitlementAccepted()
}
