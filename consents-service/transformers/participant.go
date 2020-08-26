package transformers

import (
	"errors"

	apimodels "github.com/dycons/consents/consents-service/api/models"
	datamodels "github.com/dycons/consents/consents-service/data/models"
)

// ParticipantDataToAPI contains the model-building step of the data-model-to-api-model transformer.
func ParticipantDataToAPI(dataParticipant datamodels.Participant) (*apimodels.Participant, error) {
	uuid, err := UUIDDataToAPI(dataParticipant.ID, "Participant.ID")
	if err != nil {
		message := "Transformation of Participant failed with the following errors:\n"
		return nil, errors.New(message + err.Error())
	}

	return &apimodels.Participant{
		StudyIdentifier: *uuid,
	}, nil
}
