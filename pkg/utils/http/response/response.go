package response

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/spacelavr/pandora/pkg/utils/log"
)

type Response struct {
	code int
	data interface{}
}

type Template struct {
	tpl *template.Template
}

// todo mey not here
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func Execute(files string) *Template {
	tpl, err := template.ParseFiles(files)
	return &Template{tpl: template.Must(tpl, err)}
}

// Ok returns Ok response
func Ok(data interface{}) *Response {
	return &Response{code: http.StatusOK, data: data}
}

func Created() *Response {
	return &Response{code: http.StatusCreated}
}

// NotImplemented returns not implemented response
func NotImplemented() *Error {
	return &Error{
		Code:    http.StatusNotImplemented,
		Status:  http.StatusText(http.StatusNotImplemented),
		Message: http.StatusText(http.StatusNotImplemented),
	}
}

// Unauthorized returns unauthorized response
func Unauthorized() *Error {
	return &Error{
		Code:    http.StatusUnauthorized,
		Status:  http.StatusText(http.StatusUnauthorized),
		Message: http.StatusText(http.StatusUnauthorized),
	}
}

// InvalidJSON returns invalid json response
func InvalidJSON() *Error {
	return &Error{
		Code:    http.StatusUnprocessableEntity,
		Status:  http.StatusText(http.StatusUnprocessableEntity),
		Message: "invalid json",
	}
}

// Forbidden returns forbidden response
func Forbidden() *Error {
	return &Error{
		Code:    http.StatusForbidden,
		Status:  http.StatusText(http.StatusForbidden),
		Message: http.StatusText(http.StatusForbidden),
	}
}

// NotFound returns not found response
func NotFound(what string) *Error {
	return &Error{
		Code:    http.StatusNotFound,
		Status:  http.StatusText(http.StatusNotFound),
		Message: fmt.Sprintf("%s not found", what),
	}
}

// AlreadyExists returns already exists response
func AlreadyExists(what string) *Error {
	return &Error{
		Code:    http.StatusConflict,
		Status:  http.StatusText(http.StatusConflict),
		Message: fmt.Sprintf("%s already exists", what),
	}
}

// InternalServerError returns internal server error response
func InternalServerError() *Error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Status:  http.StatusText(http.StatusInternalServerError),
		Message: http.StatusText(http.StatusInternalServerError),
	}
}

// BadParameter returns bad parameter response
func BadParameter(parameter string) *Error {
	return &Error{
		Code:    http.StatusUnprocessableEntity,
		Status:  http.StatusText(http.StatusUnprocessableEntity),
		Message: fmt.Sprintf("bad %s parameter", parameter),
	}
}

// Http send http response
func (r *Response) Http(w http.ResponseWriter) {
	w.WriteHeader(r.code)
	if err := json.NewEncoder(w).Encode(r.data); err != nil {
		log.Error(err)
	}
}

func (t *Template) Http(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
	if err := t.tpl.Execute(w, nil); err != nil {
		InternalServerError().Http(w)
	}
}

func (r *Error) Http(w http.ResponseWriter) {
	w.WriteHeader(r.Code)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Error(err)
	}
}
