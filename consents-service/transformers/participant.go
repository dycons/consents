package transformers

import (
	"github.com/go-openapi/strfmt"

	apimodels "github.com/dycons/consents/consents-service/api/models"
	datamodels "github.com/dycons/consents/consents-service/data/models"
)

// ParticipantDataToAPI contains the model-building step of the data-model-to-api-model transformer.
func ParticipantDataToAPI(dataParticipant datamodels.Participant) (*apimodels.Participant, error) {
	return &apimodels.Participant{
		StudyIdentifier: strfmt.UUID(dataParticipant.ID.String()),
	}, nil
}
