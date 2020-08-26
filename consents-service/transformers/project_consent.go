package transformers

import (
	apimodels "github.com/dycons/consents/consents-service/api/models"
	datamodels "github.com/dycons/consents/consents-service/data/models"
)

// ProjectConsentDataToAPI contains the model-building step of the data-model-to-api-model transformer.
func ProjectConsentDataToAPI(dataProjectConsent datamodels.ProjectConsent) (*apimodels.ProjectConsent, error) {
	applicationID := int32(dataProjectConsent.ProjectApplicationID)
	geneticConsent := dataProjectConsent.GeneticConsent
	clinicalConsent := dataProjectConsent.ClinicalConsent

	return &apimodels.ProjectConsent{
		ApplicationID:   &applicationID,
		GeneticConsent:  &geneticConsent,
		ClinicalConsent: &clinicalConsent}, nil
}
