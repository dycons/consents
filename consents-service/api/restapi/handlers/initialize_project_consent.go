package handlers

import (
	// TODO rm

	"fmt"

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

// InitializeProjectConsent takes the default consents for the study_identifier posted by the API request,
// and uses them to initialize a project-specific set of consents, which it creates into the database.
// It then returns the URL location of these project consents.
func InitializeProjectConsent(params operations.InitializeProjectConsentParams, tx *pop.Connection) middleware.Responder {
	// Transform the Participant UUID from the api format to the data format
	participantID, err := transformers.UUIDAPIToData(params.StudyIdentifier, "StudyIdentifier")
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Transforming the StudyIdentifier from API to data formats failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewInitializeProjectConsentInternalServerError().WithPayload(errPayload)
	}

	// Check that a ProjectConsent for this combination of application id and study id has not been initialized already
	existingProjectConsent := datamodels.ProjectConsent{}
	query := fmt.Sprintf("participant_id = '%s' AND project_application_id = '%d'",
		participantID, *params.ProjectConsentInitialization.ProjectApplicationID)
	projectConsentExists, err := tx.Where(query).Exists(&existingProjectConsent)
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Checking for duplication of ProjectConsentInitialization failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewInitializeProjectConsentInternalServerError().WithPayload(errPayload)
	}
	if projectConsentExists {
		message := "Duplicates the request for the initialization of a ProjectConsent."
		var code int64 = 403002

		log.Write(params.HTTPRequest, code, err).Warn(message)
		errPayload := &apimodels.Error{Code: &code, Message: &message}
		return operations.NewInitializeProjectConsentForbidden().WithPayload(errPayload)
	}

	// Find the DefaultConsent associated with the uuid given in the request
	defaultConsent, err := utilities.GetDefaultConsentByParticipantID(participantID.String(), tx)
	if err != nil {
		message := "This default consent cannot be found"
		var code int64 = 404001

		log.Write(params.HTTPRequest, code, err).Warn(message)
		errPayload := &apimodels.Error{Code: &code, Message: &message}
		return operations.NewInitializeProjectConsentNotFound().WithPayload(errPayload)
	}

	// Initialize the consent choise for genetic data from the overall consent
	var geneticConsent bool
	switch defaultConsent.GeneticConsentStyle {
	case datamodels.SecondaryUseForbidden:
		geneticConsent = false
	case datamodels.OptIn:
		geneticConsent = false
	case datamodels.OptOut:
		geneticConsent = true
	default:
		message := "Translation of GeneticConsentStyle into consent choices fails to yield valid enum. Got: " + string(defaultConsent.GeneticConsentStyle)
		log.Write(params.HTTPRequest, 500000, err).Error(message)
		errPayload := errors.DefaultInternalServerError()
		return operations.NewInitializeProjectConsentInternalServerError().WithPayload(errPayload)
	}

	// Initialize the consent choise for clinical data from the overall consent style
	var clinicalConsent bool
	switch defaultConsent.ClinicalConsentStyle {
	case datamodels.SecondaryUseForbidden:
		clinicalConsent = false
	case datamodels.OptIn:
		clinicalConsent = false
	case datamodels.OptOut:
		clinicalConsent = true
	default:
		message := "Translation of ClinicalConsentStyle into consent choices fails to yield valid enum. Got: " + string(defaultConsent.ClinicalConsentStyle)
		log.Write(params.HTTPRequest, 500000, err).Error(message)
		errPayload := errors.DefaultInternalServerError()
		return operations.NewInitializeProjectConsentInternalServerError().WithPayload(errPayload)
	}

	// Create the ProjectConsent into the DB. Only procees if creation succeeds.
	projectConsent := datamodels.ProjectConsent{
		ParticipantID:        *participantID,
		ProjectApplicationID: int(*params.ProjectConsentInitialization.ProjectApplicationID),
		GeneticConsent:       geneticConsent,
		ClinicalConsent:      clinicalConsent,
	}
	err = tx.Create(&projectConsent)
	if err != nil {
		log.Write(params.HTTPRequest, 500000, err).Error("Creating into database failed")
		errPayload := errors.DefaultInternalServerError()
		return operations.NewInitializeProjectConsentInternalServerError().WithPayload(errPayload)
	}

	// Return the status of the Initialization and the path to the created ProjectConsent
	status := apimodels.ProjectConsentInitializationStatusComplete
	initialization := apimodels.ProjectConsentInitialization{
		ProjectApplicationID: params.ProjectConsentInitialization.ProjectApplicationID,
		StudyIdentifier:      params.StudyIdentifier,
		Status:               &status,
	}
	location := params.HTTPRequest.URL.Host + params.HTTPRequest.URL.EscapedPath() + string(projectConsent.ProjectApplicationID)
	return operations.NewInitializeProjectConsentCreated().WithPayload(&initialization).WithLocation(location)
}
