package transformers

import (
	"github.com/go-openapi/strfmt"

	apimodels "github.com/dycons/consents/consents-service/api/models"
	datamodels "github.com/dycons/consents/consents-service/data/models"
)

// DefaultConsentAPIToData contains the model-building step of the api-model-to-data-model transformer.
// Since there is no Participant property contained in the DefaultConsent API model, the ParticipantID and
// Participant fields of the data model are not populated in this method, and must be handled seperately.
func DefaultConsentAPIToData(apiDefaultConsent apimodels.DefaultConsent) (*datamodels.DefaultConsent, error) {
	switch *apiDefaultConsent.GeneticConsentStyle {
	case apimodels.DefaultConsentGeneticConsentStyleSUF:
		dataGeneticConsentStyle := datamodels.SecondaryUseForbidden
	case apimodels.DefaultConsentGeneticConsentStyleOI:
		dataGeneticConsentStyle := datamodels.OptIn
	case apimodels.DefaultConsentGeneticConsentStyleOO:
		dataGeneticConsentStyle := datamodels.OptOut
	default:
		message := "Transformation of GeneticConsentStyle from api to data model fails to yield valid enum."
		return nil, errors.New(message)
	}

	switch *apiDefaultConsent.ClinicalConsentStyle {
	case apimodels.DefaultConsentClinicalConsentStyleSUF:
		dataClinicalConsentStyle := datamodels.SecondaryUseForbidden
	case apimodels.DefaultConsentClinicalConsentStyleOI:
		dataClinicalConsentStyle := datamodels.OptIn
	case apimodels.DefaultConsentClinicalConsentStyleOO:
		dataClinicalConsentStyle := datamodels.OptOut
	default:
		message := "Transformation of ClinicalConsentStyle from api to data model fails to yield valid enum."
		return nil, errors.New(message)
	}

	return &datamodels.DefaultConsent{
		GeneticConsentStyle:  dataGeneticConsentStyle,
		ClinicalConsentStyle: dataClinicalConsentStyle}, nil
}
