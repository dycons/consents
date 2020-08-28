package transformers

import (
	apimodels "github.com/dycons/consents/consents-service/api/models"
	datamodels "github.com/dycons/consents/consents-service/data/models"
)

// ProjectConsentDataToAPI contains the model-building step of the data-model-to-api-model transformer.
func ProjectConsentDataToAPI(dataProjectConsent datamodels.ProjectConsent) (*apimodels.ProjectConsent, error) {
	projectApplicationID := int32(dataProjectConsent.ProjectApplicationID)

	return &apimodels.ProjectConsent{
		ProjectApplicationID: &projectApplicationID,
		GeneticConsent:       &dataProjectConsent.GeneticConsent,
		ClinicalConsent:      &dataProjectConsent.ClinicalConsent}, nil
}
