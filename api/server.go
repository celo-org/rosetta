package api

import (
	"net/http"
)

// BaddRequest replies to the request with status 400 and the error message
// It does not otherwise end the request; the caller should ensure no further
func BadRequest(w http.ResponseWriter, err error) {
	payload := BuildErrorResponse(int32(http.StatusBadRequest), err)
	JSONResponse(payload, http.StatusBadRequest, w)
}
