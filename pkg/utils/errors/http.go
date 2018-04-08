package errors

import (
	"encoding/json"
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
