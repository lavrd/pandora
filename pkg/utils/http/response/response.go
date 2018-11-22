package response

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"pandora/pkg/utils/errors"
	"pandora/pkg/utils/log"
)

// Response
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Template
type Template struct {
	tpl *template.Template
}

// ContentResponse
type ContentResponse struct {
	Data        []byte
	ContentType string
}

// Execute execute template
func Execute(files string) *Template {
	tpl, err := template.ParseFiles(files)
	return &Template{tpl: template.Must(tpl, err)}
}

// AlreadyExists returns already exists response
func AlreadyExists(message string) *Response {
	return &Response{
		Code:    http.StatusConflict,
		Status:  http.StatusText(http.StatusConflict),
		Message: message,
	}
}

// BadParameter returns bad parameter response
func BadParameter(message string) *Response {
	return &Response{
		Code:    http.StatusUnprocessableEntity,
		Status:  http.StatusText(http.StatusUnprocessableEntity),
		Message: fmt.Sprintf("You passed the wrong parameter. %s", message),
	}
}

// NotFound returns not found response
func NotFound(message string) *Response {
	return &Response{
		Code:    http.StatusNotFound,
		Status:  http.StatusText(http.StatusNotFound),
		Message: message,
	}
}

// OK returns ok response
func OK() *Response {
	return &Response{
		Message: http.StatusText(http.StatusOK),
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
	}
}

// JSON returns json response
func JSON(data interface{}) *ContentResponse {
	buf, err := json.Marshal(data)
	if err != nil {
		log.Error(errors.WithStack(err))
	}

	return &ContentResponse{
		Data:        buf,
		ContentType: "application/json",
	}
}

// InternalServerError returns interval server error response
func InternalServerError() *Response {
	return &Response{
		Code:    http.StatusInternalServerError,
		Status:  http.StatusText(http.StatusInternalServerError),
		Message: "Oops, something went wrong",
	}
}

// InvalidJSON returns invalid json response
func InvalidJSON(message string) *Response {
	return &Response{
		Code:    http.StatusUnsupportedMediaType,
		Status:  http.StatusText(http.StatusUnsupportedMediaType),
		Message: fmt.Sprintf("You need one of these parameters: %s", message),
	}
}

func (r *ContentResponse) HTTP(w http.ResponseWriter) {
	w.Header().Set("Content-Type", r.ContentType)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(r.Data); err != nil {
		log.Error(errors.WithStack(err))
	}
}

// HTTP send http response
func (r *Response) HTTP(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Error(errors.WithStack(err))
	}
}

// HTTP send template
func (t *Template) HTTP(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
	if err := t.tpl.Execute(w, nil); err != nil {
		InternalServerError().HTTP(w)
	}
}
