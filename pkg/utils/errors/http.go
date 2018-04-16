package errors

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spacelavr/pandora/pkg/log"
)

// Response
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status"`
}

// NotImplemented returns not implemented response
func NotImplemented() *Response {
	return &Response{Code: http.StatusNotImplemented, Status: http.StatusText(http.StatusNotImplemented)}
}

// Unauthorized returns unauthorized response
func Unauthorized() *Response {
	return &Response{Code: http.StatusUnauthorized, Status: http.StatusText(http.StatusUnauthorized)}
}

// InvalidJSON returns invalid json response
func InvalidJSON() *Response {
	return &Response{
		Code:    http.StatusUnprocessableEntity,
		Status:  http.StatusText(http.StatusUnprocessableEntity),
		Message: "invalid json",
	}
}

// Forbidden returns forbidden response
func Forbidden() *Response {
	return &Response{Code: http.StatusForbidden, Status: http.StatusText(http.StatusForbidden)}
}

// NotFound returns not found response
func NotFound(what string) *Response {
	return &Response{
		Code:    http.StatusNotFound,
		Status:  http.StatusText(http.StatusNotFound),
		Message: fmt.Sprintf("%s not found", what),
	}
}

// AlreadyExists returns already exists response
func AlreadyExists(what string) *Response {
	return &Response{
		Code:    http.StatusConflict,
		Status:  http.StatusText(http.StatusConflict),
		Message: fmt.Sprintf("%s already exists", what),
	}
}

// InternalServerError returns internal server error response
func InternalServerError() *Response {
	return &Response{Code: http.StatusInternalServerError, Status: http.StatusText(http.StatusInternalServerError)}
}

// BadParameter returns bad parameter response
func BadParameter(parameter string) *Response {
	return &Response{
		Code:    http.StatusUnprocessableEntity,
		Status:  http.StatusText(http.StatusUnprocessableEntity),
		Message: fmt.Sprintf("bad %s parameter", parameter),
	}
}

// Http send http response
func (r *Response) Http(w http.ResponseWriter) {
	r.send(w)
}

func (r Response) send(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Error(err)
	}
}
