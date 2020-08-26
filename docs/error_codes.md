# Error Codes

This is the set of 5-digit model-vs-specific error codes returned by the variant service.

The first 3 digits of the error code are the HTTP status code that the error is returned under, with the following 3 digits being specific to this service.

The code descriptions provided here are not intended to replace the error messages returned alongside them, but rather to facilitate automation of error handling and/or logging on the client side. Therefore they may be either identical to or more descriptive than the messages returned in error payloads.

## 400--- Codes

##### 403001

Attempted a forbidden "GET all"-type request.

GET requests for this resource are permitted only with some query parameters provided in the query string.

##### 403002

Attempted to post a forbidden duplicate of an existing resource.

##### 404001

Requested record cannot be found. The given record ID does not correspond with any existing data for this resource.

##### 404002

The resource by which you are attempting to query for other data cannot be found. The given record ID does not correspond with any existing data for this resource.

## 500--- Codes

##### 500000

This is the default internal server error code. No further information can be provided here; please view the issues on the service [project repository](https://github.com/CanDIG/go-model-service/issues) and create a new one if it has not already been brought to our attention. Please include the following in your issue report:
- A concise and issue-specific Title
- Steps to reproduce
- Expected Result
- Actual Result

We thank you in advance for contributing to the improvement of this service!
