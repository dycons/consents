/*
Package utilities implements general-purpose utility functions for use by the restapi handlers.
*/
package utilities

import (
	apimodels "github.com/dycons/consents/consents-service/api/models"
	"github.com/dycons/consents/consents-service/errors"
	"github.com/dycons/consents/utilities/log"
	"github.com/gobuffalo/pop"
	"net/http"
)

// ConnectDevelopment connects to the development database and returns the connection and/or error message
func ConnectDevelopment(HTTPRequest *http.Request) (*pop.Connection, *apimodels.Error) {
	tx, err := pop.Connect("development")
	if err != nil {
		log.Write(HTTPRequest, 500000, err).Error("Failed to connect to database")
		errPayload := errors.DefaultInternalServerError()
		return nil, errPayload
	}
	return tx, nil
}
