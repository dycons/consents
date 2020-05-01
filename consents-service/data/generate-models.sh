#!/bin/bash

# Generate model for StudyParticipant.
# This entity only needs the ID field for associations, so no other fields are generated.
soda generate model study_participant -e development

# Generate model for the ConsentStyle lookup table, which represents the enum of [OI, OO, DNS] (opt-in, opt-out, do-not-share)
# but can also be extended to other consent styles more easily than actual enums.
# This entity is generally loaded eagerly when the DefaultConsent entity is loaded
soda generate model consent_style \
	consent_style:string \
	-e development

# Generate model for DefaultConsent
# TODO associations, enums
soda generate model default_consent \
	study_participant_id:uuid \
	genetic_consent_style_id:uuid \
	clinical_consent_style_id:uuid \
	-e development

# Generate model for ProjectConsent
# TODO associations, check if bool is correct type
soda generate model project_consent \
	study_participant_id:uuid \
	project_application_id:uuid \
	genetic_consent:bool \
	clinical_consent:bool \
	-e development
