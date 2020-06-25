#!/bin/bash

# Generate model for StudyParticipant.
# This entity only needs the ID field for associations, so no other fields are generated.
soda generate model study_participant -e development

# Generate model for DefaultConsent
# All three associations below should have ForeignKeys added to them
soda generate model default_consent \
	study_participant_id:uuid \
	genetic_consent_style:int \
	clinical_consent_style:int \
	-e development

# Generate model for ProjectConsent
# StudyParticipant association should have ForeignKeys added to it
soda generate model project_consent \
	study_participant_id:uuid \
	project_application_id:uuid \
	genetic_consent:bool \
	clinical_consent:bool \
	-e development
