package transformers

import (
	"errors"

	apimodels "github.com/dycons/consents/consents-service/api/models"
	datamodels "github.com/dycons/consents/consents-service/data/models"
	"github.com/gobuffalo/uuid"
)

// DefaultConsentAPIToData contains the model-building step of the api-model-to-data-model transformer.
// Since there is no Participant property contained in the DefaultConsent API model, the ParticipantID and
// Participant fields of the data model are not populated in this method, and must be handled seperately.
func DefaultConsentAPIToData(apiDefaultConsent apimodels.DefaultConsent, participantID uuid.UUID) (*datamodels.DefaultConsent, error) {
	var dataGeneticConsentStyle int
	switch apiDefaultConsent.GeneticConsentStyle {
	case apimodels.DefaultConsentGeneticConsentStyleSUF:
		dataGeneticConsentStyle = datamodels.SecondaryUseForbidden
	case apimodels.DefaultConsentGeneticConsentStyleOI:
		dataGeneticConsentStyle = datamodels.OptIn
	case apimodels.DefaultConsentGeneticConsentStyleOO:
		dataGeneticConsentStyle = datamodels.OptOut
	default:
		message := "Transformation of GeneticConsentStyle from api to data model fails to yield valid enum. Got: " + apiDefaultConsent.GeneticConsentStyle
		return nil, errors.New(message)
	}

	var dataClinicalConsentStyle int
	switch apiDefaultConsent.ClinicalConsentStyle {
	case apimodels.DefaultConsentClinicalConsentStyleSUF:
		dataClinicalConsentStyle = datamodels.SecondaryUseForbidden
	case apimodels.DefaultConsentClinicalConsentStyleOI:
		dataClinicalConsentStyle = datamodels.OptIn
	case apimodels.DefaultConsentClinicalConsentStyleOO:
		dataClinicalConsentStyle = datamodels.OptOut
	default:
		message := "Transformation of ClinicalConsentStyle from api to data model fails to yield valid enum. Got: " + apiDefaultConsent.ClinicalConsentStyle
		return nil, errors.New(message)
	}

	return &datamodels.DefaultConsent{
		ParticipantID:        participantID,
		GeneticConsentStyle:  dataGeneticConsentStyle,
		ClinicalConsentStyle: dataClinicalConsentStyle}, nil
}

// DefaultConsentDataToAPI contains the model-building step of the data-model-to-api-model transformer.
func DefaultConsentDataToAPI(dataDefaultConsent datamodels.DefaultConsent) (*apimodels.DefaultConsent, error) {
	var apiGeneticConsentStyle string
	switch dataDefaultConsent.GeneticConsentStyle {
	case datamodels.SecondaryUseForbidden:
		apiGeneticConsentStyle = apimodels.DefaultConsentGeneticConsentStyleSUF
	case datamodels.OptIn:
		apiGeneticConsentStyle = apimodels.DefaultConsentGeneticConsentStyleOI
	case datamodels.OptOut:
		apiGeneticConsentStyle = apimodels.DefaultConsentGeneticConsentStyleOO
	default:
		message := "Transformation of GeneticConsentStyle from data to api model fails to yield valid enum. Got: " + string(dataDefaultConsent.GeneticConsentStyle)
		return nil, errors.New(message)
	}

	var apiClinicalConsentStyle string
	switch dataDefaultConsent.ClinicalConsentStyle {
	case datamodels.SecondaryUseForbidden:
		apiClinicalConsentStyle = apimodels.DefaultConsentClinicalConsentStyleSUF
	case datamodels.OptIn:
		apiClinicalConsentStyle = apimodels.DefaultConsentClinicalConsentStyleOI
	case datamodels.OptOut:
		apiClinicalConsentStyle = apimodels.DefaultConsentClinicalConsentStyleOO
	default:
		message := "Transformation of ClinicalConsentStyle from data to api model fails to yield valid enum. Got: " + string(dataDefaultConsent.ClinicalConsentStyle)
		return nil, errors.New(message)
	}

	return &apimodels.DefaultConsent{
		GeneticConsentStyle:  apiGeneticConsentStyle,
		ClinicalConsentStyle: apiClinicalConsentStyle}, nil
}
