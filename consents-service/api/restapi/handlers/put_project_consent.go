package handlers

import (
	"fmt"

	apimodels "github.com/dycons/consents/consents-service/api/models"
	"github.com/dycons/consents/consents-service/api/restapi/operations"
	datamodels "github.com/dycons/consents/consents-service/data/models"
	"github.com/dycons/consents/consents-service/errors"
	"github.com/dycons/consents/consents-service/transformers"
	"github.com/dycons/consents/consents-service/utilities/log"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gobuffalo/pop"
)

// PutProjectConsent takes the default consents for the study_identifier posted by the API request,
// and uses them to initialize a project-specific set of consents, which it creates into the database.
// It then returns the URL location of these project consents.
func PutProjectConsent(params operations.PutProjectConsentParams, tx *pop.Connection) middleware.Responder {
	// Transform the Participant UUID from the api format to the data format
	participantID, err := transformers.UUIDAPIToData(params.StudyIdentifier, "StudyIdentifier")
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Transforming the StudyIdentifier from API to data formats failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewInitializeProjectConsentInternalServerError().WithPayload(errPayload)
	}

	// Fetch the ProjectConsent for this combination of project application and participant
	dataProjectConsent := datamodels.ProjectConsent{}
	query := fmt.Sprintf("participant_id = '%s' AND project_application_id = '%d'",
		participantID, *params.ProjectConsent.ProjectApplicationID)
	err = tx.Where(query).First(&dataProjectConsent)
	if err != nil {
		// This ProjectConsent does not exist in the db; return a 404 Not Found response
		if err.Error() == "sql: no rows in result set" {
			message := "This ProjectConsent cannot be found"
			var code int64 = 404001

			log.Write(params.HTTPRequest, code, err).Warn(message)
			errPayload := &apimodels.Error{Code: &code, Message: &message}
			return operations.NewPutProjectConsentNotFound().WithPayload(errPayload)
		}
		// Handle any other error as an internal server error
		log.Write(params.HTTPRequest, 500000, err).Error("Retrieving the existing ProjectConsent from the database failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewInitializeProjectConsentInternalServerError().WithPayload(errPayload)
	}

	// Update the ProjectConsent with the data in the payload
	dataProjectConsent.GeneticConsent = *params.ProjectConsent.GeneticConsent
	dataProjectConsent.ClinicalConsent = *params.ProjectConsent.ClinicalConsent

	// TODO Validate the dataProjectConsent and then Update into the db
	// TODO Make sure that all of my other Create() calls are being preceded by calls to Validate(), or to ValidateAndCreate()

	// Write the ProjectConsent into the DB. Only procees if write succeeds.
	validationErrors, err := tx.ValidateAndUpdate(&dataProjectConsent)
	if validationErrors.Error() != "" { // if at least one validation error occured
		log.Write(params.HTTPRequest, 500000, err).Error("Data schema validation for the ProjectConsent failed with the following validation errors: " +
			validationErrors.Error())
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPutProjectConsentInternalServerError().WithPayload(errPayload)
	}
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Creation of the ProjectConsent into the database failed without validation errors.")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewPutProjectConsentInternalServerError().WithPayload(errPayload)
	}

	// Return the updated ProjectConsent
	updatedProjectConsent, errPayload := projectConsentDataToAPIModel(dataProjectConsent, params.HTTPRequest)
	if errPayload != nil {
		return operations.NewPostParticipantInternalServerError().WithPayload(errPayload)
	}
	return operations.NewPutProjectConsentOK().WithPayload(updatedProjectConsent)
}
