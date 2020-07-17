package transformers

import (
	"github.com/go-openapi/strfmt"

	apimodels "github.com/dycons/consents/consents-service/api/models"
	datamodels "github.com/dycons/consents/consents-service/data/models"
)

// ParticipantDataToAPI contains the model-building step of the data-model-to-api-model transformer.
// Presently, this transformer always populates the DefaultConsent property with a nil pointer, as this property
// is not expected in any of the responses in the v1 API spec.
func ParticipantDataToAPI(dataParticipant datamodels.Participant) (*apimodels.Participant, error) {
	return &apimodels.Participant{
		ID:             strfmt.UUID(dataParticipant.ID.String()),
		DefaultConsent: nil}, nil
}
