/*
Package errors implements error reporting functionality that is commonly used within this project.
Its purpose is to standardize generic error reports, both for external use via API responses and for
internal use via logging.
*/
package errors

import (
	apimodels "github.com/dycons/consents/consents-service/api/models"
	"github.com/go-openapi/swag"
)

// DefaultInternalServerError returns an error-non-specific payload for a generic 500 server response.
// For the sake of both simplicity and security, there should not be any further detail returned to the
// client in a 500: Internal Server Error response.
func DefaultInternalServerError() *apimodels.Error {
	var code int64 = 500000
	return &apimodels.Error{Code: &code, Message: swag.String("An internal server error has occurred")}
}
