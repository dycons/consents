/*
Package utilities implements general-purpose utility functions for use by the restapi handlers.
*/
package utilities

import (
	"errors"

	datamodels "github.com/dycons/consents/consents-service/data/models"
)

// InitializeProjectConsents initializes a geneticconsent and a clinical consent based on the
// default consent preferences.
// DefaultConsent options are interpreted as follows:
// 			SecondaryUseForbidden 	-> false (do not consent by default)
//			OptIn 					-> false (do not consent by default)
// 			OptOut 					-> true (consent by default)
func InitializeProjectConsents(defaultConsent datamodels.DefaultConsent) (geneticConsent bool, clinicalConsent bool, err error) {
	// Initialize the consent choise for genetic data from the overall consent
	switch defaultConsent.GeneticConsentStyle {
	case datamodels.SecondaryUseForbidden:
		geneticConsent = false
	case datamodels.OptIn:
		geneticConsent = false
	case datamodels.OptOut:
		geneticConsent = true
	default:
		message := "Translation of GeneticConsentStyle into consent choices fails to yield valid enum. Got: " + string(defaultConsent.GeneticConsentStyle)
		return false, false, errors.New(message)
	}

	// Initialize the consent choise for clinical data from the overall consent style
	switch defaultConsent.ClinicalConsentStyle {
	case datamodels.SecondaryUseForbidden:
		clinicalConsent = false
	case datamodels.OptIn:
		clinicalConsent = false
	case datamodels.OptOut:
		clinicalConsent = true
	default:
		message := "Translation of GeneticConsentStyle into consent choices fails to yield valid enum. Got: " + string(defaultConsent.GeneticConsentStyle)
		return false, false, errors.New(message)
	}

	return geneticConsent, clinicalConsent, nil
}
